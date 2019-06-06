package main

import (
    "fmt"
    "log"

    "github.com/llir/llvm/asm"
    "github.com/llir/llvm/ir"
    "github.com/llir/llvm/ir/metadata"
)

func main() {
    // Parse the target/goal ll we're trying to get to
    tar, err := asm.ParseFile("target.ll")
    if err != nil {
        log.Fatalf("%+v", err)
    }

    // Parse the ll we're working with currently
    wor, err := asm.ParseFile("foo.ll")
    if err != nil {
        log.Fatalf("%+v", err)
    }

    // TODO: Display the differences between the two ll files
    //   asm/module.go has some decent pieces which *might* be usable, but they're not exported :/
    // TODO: Once this is working ok(?), reduce the output to only display meaningful differences
    fmt.Println("Differences to investigate:")
    if tar.SourceFilename != wor.SourceFilename {
        // TODO: This difference can probably be skipped
        fmt.Println("  * Source filename")
        fmt.Printf("      Target: '%s'\n", tar.SourceFilename)
        fmt.Printf("        Work: '%s'\n\n", wor.SourceFilename)
    }
    if tar.DataLayout != wor.DataLayout {
        fmt.Println("  * Data layout")
        fmt.Printf("      Target: '%s'\n", tar.DataLayout)
        fmt.Printf("        Work: '%s'\n\n", wor.DataLayout)
    }
    if tar.TargetTriple != wor.TargetTriple {
        fmt.Println("  * Target triple")
        fmt.Printf("      Target: '%s'\n", tar.TargetTriple)
        fmt.Printf("        Work: '%s'\n", wor.TargetTriple)
    }
    fmt.Println()

    // * Check the basic info of each function *

    // Compare the # of functions
    if len(tar.Funcs) != len(wor.Funcs){
        fmt.Println("  * Number of functions")
        fmt.Printf("      Target: '%d'\n", len(tar.Funcs))
        fmt.Printf("        Work: '%d'\n", len(wor.Funcs))
    }

    // Use a map to hold the functions being compared, as the function ordering could be different between the two ll files
    worFuncs := make(map[string]*ir.Func)
    for _, j := range wor.Funcs {
        worFuncs[j.GlobalName] = j
    }
    for _, j := range tar.Funcs {
        nm := j.GlobalIdent.GlobalName
        noDiffs := true

        // Nothing to check for this function if it's not present in the working ll
        wf, ok := worFuncs[nm]
        if !ok {
            fmt.Printf("  * Function '%s' is only present in target\n", nm)
            continue
        }
        fmt.Printf("  * Function '%s'\n", nm)

        // Compare the return types
        if !j.Sig.RetType.Equal(wf.Sig.RetType) {
            noDiffs = false
            fmt.Println("    * Return type")
            fmt.Printf("        Target: '%v'\n", j.Sig.RetType)
            fmt.Printf("          Work: '%v'\n", wf.Sig.RetType)
        }

        // Compare the # of parameters
        if len(j.Params) != len(wf.Params) {
            noDiffs = false
            fmt.Println("    * # of Parameters")
            fmt.Printf("        Target: '%d'\n", len(j.Params))
            fmt.Printf("          Work: '%d'\n", len(wf.Params))
        }

        // Compare the type of each parameter
        for k, l := range j.Params {
            if !l.Typ.Equal(wf.Params[k].Typ) {
                noDiffs = false
                fmt.Printf("    * Type of parameter %d\n", k)
                fmt.Printf("        Target: '%s'\n", l.Typ)
                fmt.Printf("          Work: '%s'\n", wf.Params[k].Typ)
            }
        }

        // Compare the # of function attributes
        // TODO: This could be improved to compare the attributes themselves, as well as look inside ir.AttrGroupDef containers
        if len(j.FuncAttrs) != len(wf.FuncAttrs) {
            noDiffs = false
            fmt.Println("    * # of Attributes")
            fmt.Printf("        Target: '%d'\n", len(j.FuncAttrs))
            fmt.Printf("          Work: '%d'\n", len(wf.FuncAttrs))
        }

        // Purely to help with visual formatting
        if !noDiffs {
            fmt.Println()
        }
    }

    // TODO: Add check for functions only present in the working ll?

    // * Compare the _unnamed_ metadata definitions *

    // Make sure they both have the same # of metadata definitions
    if len(tar.MetadataDefs) != len(wor.MetadataDefs){
        fmt.Println("  * Number of metadata definitions")
        fmt.Printf("      Target: '%d'\n", len(tar.MetadataDefs))
        fmt.Printf("        Work: '%d'\n", len(wor.MetadataDefs))
    }

    // Ensure each of the unnamed metadata definitions present in the target ll is present in the working ll
    worDefs := make(map[string]metadata.Definition)
    for _, j := range wor.MetadataDefs {
        worDefs[j.String()] = j
    }

    for _, j := range tar.MetadataDefs {
        nm := j.String()
        wd := worDefs[nm]

        if j.LLString() != wd.LLString() {
            fmt.Printf("  * Metadata element %s is different\n", nm)
            fmt.Printf("      Target: '%s'\n", j.LLString())
            fmt.Printf("        Work: '%s'\n", wd.LLString())
        }
    }
    fmt.Println()

    // * Compare the _named_ metadata definitions *

    // Make sure they both have the same # of unnamed metadata definitions
    if len(tar.NamedMetadataDefs) != len(wor.NamedMetadataDefs){
        fmt.Println("  * Number of named Metadata definitions")
        fmt.Printf("      Target: '%d'\n", len(tar.NamedMetadataDefs))
        fmt.Printf("        Work: '%d'\n", len(wor.NamedMetadataDefs))
    }

    // Ensure each of the named metadata definitions present in the target ll is present in the working ll
    worNamedDefs := make(map[string]*metadata.NamedDef)
    for _, j := range wor.NamedMetadataDefs {
        worNamedDefs[j.Name] = j
    }

    for _, j := range tar.NamedMetadataDefs {
        nm := j.Name
        noDiffs := true

        // Nothing to check for this metadata definition if it's not present in the working ll
        nmd, ok := worNamedDefs[nm]
        if !ok {
            fmt.Printf("  * Named metadata definition '%s' is only present in target\n", nm)
            continue
        }
        fmt.Printf("  * Named metadata definition '%s'...", nm)

        // Check if the # of nodes is the same for each
        if len(j.Nodes) != len(nmd.Nodes) {
            noDiffs = false
            fmt.Printf("\n    * # of Nodes")
            fmt.Printf("        Target: '%d'\n", len(j.Nodes))
            fmt.Printf("          Work: '%d'\n", len(nmd.Nodes))
        }

        // Check each of the referenced nodes match
        for k, l := range j.Nodes {
            nodeName := l.Ident()
            z := nmd.Nodes[k].Ident()
            if nodeName != z {
                noDiffs = false
                fmt.Printf("\n    * Node '%d' differs\n", k)
                fmt.Printf("        Target: '%s'\n", nodeName)
                fmt.Printf("          Work: '%s'\n", z)
            }
        }

        if noDiffs {
            fmt.Printf(" no differences detected\n")
        } else {
            fmt.Println()
        }
    }
}

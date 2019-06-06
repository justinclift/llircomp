package main

import (
    "fmt"
    "log"

    "github.com/llir/llvm/asm"
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
        fmt.Println("  * Source filename")
        fmt.Printf("      Target: '%s'\n", tar.SourceFilename)
        fmt.Printf("        Work: '%s'\n", wor.SourceFilename)
    }
    if tar.DataLayout != wor.DataLayout {
        fmt.Println("  * Data layout")
        fmt.Printf("      Target: '%s'\n", tar.DataLayout)
        fmt.Printf("        Work: '%s'\n", wor.DataLayout)
    }
    if tar.TargetTriple != wor.TargetTriple {
        fmt.Println("  * Target triple")
        fmt.Printf("      Target: '%s'\n", tar.TargetTriple)
        fmt.Printf("        Work: '%s'\n", wor.TargetTriple)
    }

    // TODO: This will need to check each attribute, of each function
    // for i, j := range tar.Funcs {
    //     if j.GlobalName != cur.Funcs[i].GlobalName {
    //         diff.Funcs
    //     }
    // }

    // switch foo {
    // case *ast.ModuleAsm:
    //     asm := unquote(entity.Asm().Text())
    //     gen.m.ModuleAsms = append(gen.m.ModuleAsms, asm)
    // case *ast.TypeDef:
    //     ident := localIdent(entity.Name())
    //     name := getTypeName(ident)
    //     if prev, ok := gen.old.typeDefs[name]; ok {
    //         if _, ok := prev.Typ().(*ast.OpaqueType); !ok {
    //             return errors.Errorf("type identifier %q already present; prev `%s`, new `%s`", enc.Local(name), text(prev), text(entity))
    //         }
    //     }
    //     gen.old.typeDefs[name] = entity
    // case *ast.ComdatDef:
    //     name := comdatName(entity.Name())
    //     if prev, ok := gen.old.comdatDefs[name]; ok {
    //         return errors.Errorf("comdat name %q already present; prev `%s`, new `%s`", enc.Comdat(name), text(prev), text(entity))
    //     }
    //     gen.old.comdatDefs[name] = entity
    // case *ast.GlobalDecl:
    //     ident := globalIdent(entity.Name())
    //     if prev, ok := gen.old.globals[ident]; ok {
    //         return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
    //     }
    //     gen.old.globals[ident] = entity
    //     gen.old.globalOrder = append(gen.old.globalOrder, ident)
    // case *ast.IndirectSymbolDef:
    //     ident := globalIdent(entity.Name())
    //     if prev, ok := gen.old.globals[ident]; ok {
    //         return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
    //     }
    //     gen.old.globals[ident] = entity
    //     gen.old.globalOrder = append(gen.old.globalOrder, ident)
    // case *ast.FuncDecl:
    //     ident := globalIdent(entity.Header().Name())
    //     if prev, ok := gen.old.globals[ident]; ok {
    //         return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
    //     }
    //     gen.old.globals[ident] = entity
    //     gen.old.globalOrder = append(gen.old.globalOrder, ident)
    // case *ast.FuncDef:
    //     ident := globalIdent(entity.Header().Name())
    //     if prev, ok := gen.old.globals[ident]; ok {
    //         return errors.Errorf("global identifier %q already present; prev `%s`, new `%s`", ident.Ident(), text(prev), text(entity))
    //     }
    //     gen.old.globals[ident] = entity
    //     gen.old.globalOrder = append(gen.old.globalOrder, ident)
    // case *ast.AttrGroupDef:
    //     id := attrGroupID(entity.ID())
    //     if prev, ok := gen.old.attrGroupDefs[id]; ok {
    //         return errors.Errorf("attribute group ID %q already present; prev `%s`, new `%s`", enc.AttrGroupID(id), text(prev), text(entity))
    //     }
    //     gen.old.attrGroupDefs[id] = entity
    // case *ast.NamedMetadataDef:
    //     name := metadataName(entity.Name())
    //     // Multiple named metadata definitions of the same name are allowed.
    //     // They are merged into a single named metadata definition with the
    //     // nodes of each definition appended.
    //     gen.old.namedMetadataDefs[name] = append(gen.old.namedMetadataDefs[name], entity)
    // case *ast.MetadataDef:
    //     id := metadataID(entity.ID())
    //     if prev, ok := gen.old.metadataDefs[id]; ok {
    //         return errors.Errorf("metadata ID %q already present; prev `%s`, new `%s`", enc.MetadataID(id), text(prev), text(entity))
    //     }
    //     gen.old.metadataDefs[id] = entity
    // case *ast.UseListOrder:
    //     gen.old.useListOrders = append(gen.old.useListOrders, entity)
    // case *ast.UseListOrderBB:
    //     gen.old.useListOrderBBs = append(gen.old.useListOrderBBs, entity)
    //
    // }
}

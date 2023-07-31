// This file is generated by generate-std.joke script. Do not edit manually!

package git

import (
	"fmt"
	. "github.com/candid82/joker/core"
	"os"
)

func InternsOrThunks() {
	if VerbosityLevel > 0 {
		fmt.Fprintln(os.Stderr, "Lazily running slow version of git.InternsOrThunks().")
	}
	gitNamespace.ResetMeta(MakeMeta(nil, `Provides API for accessing and manipulating git repositories.

         Example:

         user=> (def db (joker.bolt/open "bolt.db" 0600))
         #'user/db
         user=> (joker.bolt/create-bucket db "users")
         nil
         user=> (def id (joker.bolt/next-sequence db "users"))
         #'user/id
         user=> id
         1
         user=> (joker.bolt/put db "users" (str id) (joker.json/write-string {:id id :name "Joe Black"}))
         nil
         user=> (joker.json/read-string (joker.bolt/get db "users" (str id)))
         {"id" 1, "name" "Joe Black"}`, "1.0"))

	gitNamespace.InternVar("config", config_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("repo"))),
			`Returns git repo's config`, "1.3"))

	gitNamespace.InternVar("head", head_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("repo"))),
			`Returns the reference where HEAD is pointing to.`, "1.3"))

	gitNamespace.InternVar("log", log_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("repo"), MakeSymbol("opts"))),
			`Returns the commit historyk from the given opts.`, "1.3"))

	gitNamespace.InternVar("object", object_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("repo"), MakeSymbol("hash"))),
			`Returns an Object with the given hash.`, "1.3"))

	gitNamespace.InternVar("open", open_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("path"))),
			`Opens a git repository from the given path. It detects if the
   repository is bare or a normal one. Throws an error if the path doesn't contain a valid repository.`, "1.3").Plus(MakeKeyword("tag"), String{S: "GitRepo"}))

	gitNamespace.InternVar("ref", ref_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("repo"), MakeSymbol("name"), MakeSymbol("resolved"))),
			`Returns the reference for a given reference name. If resolved is
   true, any symbolic reference will be resolved.`, "1.3"))

	gitNamespace.InternVar("resolve-revision", resolve_revision_,
		MakeMeta(
			NewListFrom(NewVectorFrom(MakeSymbol("repo"), MakeSymbol("revision"))),
			`Resolves revision to corresponding hash. It will always
   resolve to a commit hash, not a tree or annotated tag.`, "1.3").Plus(MakeKeyword("tag"), String{S: "String"}))

}

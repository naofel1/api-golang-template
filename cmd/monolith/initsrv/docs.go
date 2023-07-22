// Package initsrv will contain all initialization of the
// required service, repository and handler.
package initsrv

/*
	Angus: I see what you're trying to do with the initsrv package, but I think it may make the bootstrapping process
	less clear overall. For example, in [InitRepository], you instantiate all the repos and collect them in a single
	[Repository] struct. But then [InitService] immediately unpacks them into the individual services. It takes a little
	work for me to understand why these intermediate objects exist and to follow the dependency graph.

	The cmd package typically consists of a single main.go file for each binary, and those main functions that call into
	the libraries you define in internal. For small applications, bootstrapping in main is ugly but clear. For larger
	applications, have a look into dependency injection frameworks like [Wire](https://github.com/google/wire) or
	[fx](https://github.com/uber-go/fx). I think they will do what you're looking for.

	Side note: Be careful with pluralisation of your type names. In initsrv, [Repository] refers to an object containing
	many repositories. A clearer name would be [Repositories] :)
*/

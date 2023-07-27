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


	Naofel: I appreciate your insights. You're correct, the initsrv package does need refining. As part of my work,
	we manage numerous microservices along with a monolithic system that encompasses all these microservices. Prior to
	creating this package, I had all initializations crammed into the main.go file, which rendered the process rather
	chaotic and opaque. Thus, my decision to compartmentalize the initialization steps into a dedicated package.
	But I think wire package can be great for my problem, I try to check how it work and how to use it but I didn't
	find concret example. Do you have a concret example of how to use it ? Thank you again for your feedback.

	Oops, yes Repositories is better :)
*/

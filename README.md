# Subcubic Parallelized Matrix Inversion with Go
An implementation of an algorithm blockwise inversion and Strassen's matrix multiplcation algorithm to invert a matrix in subcubic time. In the hopes of optimizing said algorithm, 
we will utilizing Go's "goroutines" to parallelize independent processes.

To learn more about the algorithm and the methods of implementation, head on over to the [documentation repository.](https://github.com/SubcubicInversion/documentation)

## Contributing
To contribute to this project, make sure you've installed the latest version of Go.

Then, if there are still cards in our [Trello's](https://trello.com/b/7TIapcbJ/subcubic-inversion) to-do, jump on one of them and start making pull requests.

## Running
If you want to run the code, make sure you have the latest version of Go, run `git clone` this repository, `make` in the directory you've cloned to, and run the `subcinv` binary.

# compare_directories

Comparing 2 directories and tells all the differences

It also tells the differences between the same named files

```
compare_directories -d1 directory1 -d2 directory2

Both --dir1 and --dir2 flags are required.
Usage of compare_directories:
  -d1 string
    	the first directory to compare (shorthand for -dir1)
  -d2 string
    	the first directory to compare (shorthand for -dir2)
  -dir1 string
    	the first directory to compare (shorthand: -d1)
  -dir2 string
    	the first directory to compare (shorthand: -d1)
  -h	show this help message (shorthand for --help)
  -help
    	show this help message (shorthand: -h)
```

To install, run:

```
go install github.com/basdemir/compare_directories
```

If go is not installed on a platform supprts homebrew

```
brew install go
```




# File Summarizer CLI 

This is a starter project written in Go that will help you summarize files, folders, and their contents.

Here are a few things it can do:
- When given a file, it can tell you total size, file type, and certain things (like how many lines etc.) if applicable
- When given a folder with no sub directories, it can tell you how many files, and of which type are contained within
- If given a folder with sub directories, it can give you the approx. file size of all the contents of the subdirectory

# Usage and examples

The two main sub arguments are `file` and `dir`, which will summarize a file and directory respectively. The next keyword that follows will be the name of the said file or directory.

for directories there is a `recursive` boolean flag that can optionally try to summarize all contained sub directories. 

> [!IMPORTANT]
> The recursive flag can potnetially take a long time if you have a large file system!

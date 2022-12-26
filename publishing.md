## Steps to create eBook 
The idea is to take this repository at a point in time and self-publish. 

All of the content is written in markdown, the plan will be intitially to have an eBook but also a printed version. The ideal solution will involve getting these markdown files to PDF, EPUB and MOBI formats. Whilst also leveraging Microsoft Word as a way to add additional front cover and contents pages. 

### Grammar and Spelling check 
Before this process takes place each file will need to be confirmed that all spelling and grammar is correct, this should be done for each markdown file and uploaded to the repository before moving on. 

### Installing the tools 
I am using WSL2 on my Windows machine but these will also work on Linux based systems. 

```
sudo apt-get update 
sudo apt-get install pandoc
sudo apt-get install texlive-xetex
sudo apt-get install texlive-latex-base
sudo apt-get install texlive-fonts-recommended
sudo apt-get install texlive-fonts-extra
sudo apt-get install texlive-latex-extra
```

### Markdown to Microsoft Word 
When spelling and Grammar is complete we can make a word document / google docs which can then have further Grammar and wording checks. 

You will need to navigate your terminal into the folder location where your markdown files are stored. 

`pandoc --output test.docx --from gfm -t docx *.md`

When you have the converted file we will want to add 
- Front Cover 
- Possibly look at either using readme.md and the blog to introduce 
- A table of contents
- Special Thanks & Contributers 
- Repository Info 

### Markdown to PDF 
`pandoc --output test.pdf --from gfm --pdf-engine=xelatex --to latex *.md` 

### Resources 
- [Casey Watts - Self Publishing](https://gist.github.com/caseywatts/3d8150fe04e0d8462cfc4d51b9856d39#:~:text=Markdown%20to%20eBook,the%20kindlegen%20command%20line%20command.)
- [Kindle Direct Publishing](https://kdp.amazon.com/en_US/)
- [LeanPub](https://leanpub.com/create/book)
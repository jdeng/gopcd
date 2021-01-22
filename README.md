# gopcd
A golang reader for KODAK PCD format. https://en.wikipedia.org/wiki/Photo_CD

This is sort of a rewrite based on pcdtojpg https://pcdtojpeg.sourceforge.io/Home.html (which is in C++) but not with the full features.

# note
This wasn't an intended project. It started with a few 30 year old Photo CD ISOs randomly downloaded from archive.org. It'd be quite cumbersome to mount the ISO files, extract the PCDs, and convert them to jpeg using some bash script. As extracting the base image of the PCD format is pretty straightforward, I decided to write some golang code to do the job. Most of the photos were not interesting (they were scanned in 1993) and the default resolution is low and color is a bit off. I haven't looked at the photos since, but didn't want to just abandoned the code, so decoding higher resolution images and a custom YCC to RGB conversion were added just for completeness.

It is probably only useful for archivists. 

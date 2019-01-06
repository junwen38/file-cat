import os

#源文件
sourceDir = "/Volumes/PersonalFiles/Home2/Downloads/video/2/"
#目标文件
targetFile = "/Volumes/PersonalFiles/Home2/Downloads/target2.ts"

with open(targetFile, "ab") as fpTarget:
        files = os.listdir(sourceDir)
        for file in files:
            if file.startswith("#"):
                continue
            filename = file.split("/")[-1]
            with open(os.path.join(sourceDir, filename), "rb") as fpSource:
                fpTarget.write(fpSource.read())
            print(filename + " finished")

                
        

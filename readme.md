# About Repository

Creating gui application with golang something nightmare before sometime[ until i found sciter sdk wrapper for Golang] for me.  There is very less amount of resource available for go-sciter. So I am sharing this examples to give some contribution to create resource.

## About Programs

No one is perfect in this world, so how one can create perfect things. Program written by me may be written in some another[ far better ] way. If you found something like that please make a pull request. I will be happy to review and merge it.
I have explained almost all program listed below on my blog  [ link to my blog ](https://www.mchampaneri.in)

## List of programs 

#### [01- HelloSciter](https://www.mchampaneri.in/2018/07/hello-sciter-program.html)
Hello world equivilate program for sciter. Just to get you excited by showing window on screen.

#### 02-HelloTIScript
TIScript is  extended version of JavaScript. This program contains small portion of TIScript to introduce TIScript. 

#### [03-TIScriptInput]((https://www.mchampaneri.in/2018/07/first-program-with-tiscript-and-sciter.html))
Moving one step forward in TIScript. This program shows how to access data from html inputs using TIScript.

#### [04-CallGoFunctionFromTIscript](https://www.mchampaneri.in/2018/07/process-input-grabed-from-tiscript-in.html)
Once you get input from HTML elements you may need to call a goLang function and process input. This program is simplest example of taking input from html via TIScript passing that data to goLang function and updating output on HTML element.

#### [05-Calc[EndOfPart1]](https://www.mchampaneri.in/2018/07/simple-calc-using-golang-and-sciter-sdk.html)
Summation of the journey until.  It's a very simple calc which uses every thing we learn in earlier examples.

#### [06-BuiltinHTML](https://www.mchampaneri.in/2018/07/embed-gui-inside-your-go-code.html)
It's not good if you have to load external html file. As you have to care about location of html file. But instead what if html is inside your gocode? . It's cool right!. This is examples is about how to embedded html in your go code.

#### [07-notepadScratch](https://www.mchampaneri.in/2018/07/simple-documnet-based-appliaction-with.html)
Time to explore something new. This example contains code for extremely simple notepad. It just allows you to open file, write a new file , save it and exit the appliaction. But, Its a good example for those who want to make an application has to work with documents, right! of course.

#### [08-packfolderIntro](https://www.mchampaneri.in/2018/08/use-packfolder-to-archive-your-resource.html)
If you have gone through 06, it actully puts html/css code in go file. Which looks weird. Sciter-sdk comes with one utility called packfolder which can generate single archive for your resource folder in one of the supported output format. So, now there is no need to write your gui inside go code. You can use packfolder to use that code. 

If you are confused what i am talking about, please see the code, it might make you more clear.

#### 09-image-viewer
Image-viewer support png/jpg file to view. It autoloads every jpg/png file behind the scene and displays on screen existing in the same folder as executable.
![Image of Image-Viewer](https://github.com/mchampaneri/go-sciter-example/blob/master/09-image-viewer/image-viewer.png)
 Its image loading logic is written in golang while front-end is sciter.  UI may be create even better, but as this is just for example. 

#### 10-screen-sefli
Screen-sefli takes snapshot of screen according to provided cordinates.
![Image of Screen-selfi](https://github.com/mchampaneri/go-sciter-example/blob/master/10-screen-selfi/selfi-sefli.png)

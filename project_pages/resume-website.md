# Resume Website

If you want to view the repository <a href=https://github.com/colorfulparadox/resume-website target=_blank>click here</a>

### Technology Stack and Motivation
For this project, I wanted to keep it quite simple. I also wanted to explore the other parts of development that are not usually covered in class. Such as learning to use Docker, hosting, and serverside rendering.

Everything on this website is generated on the server using Templ. Templ is a Go-lang templating library. The website is hosted on fly.io inside of a Docker image. The program that is being run in the Docker image is just a simple HTTP server. For the frontend, I just have plain HTML and CSS.

The one interesting thing I did for this project was that all the pages in my “Portfolio” section are generated from Markdown files. For this, I used a Go-lang library called Goldmark. From the generated HTML I then am able to insert it into a templ component.

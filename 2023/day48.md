# Day 47 - Let's build an App in Python

Let's create a simple blog app with the help of [Flask](https://flask.palletsprojects.com/en/2.2.x/) that supports posts in [markdown.](https://www.markdownguide.org/basic-syntax/)

## Initiating virtual env and installing packages

Let's create a directory for our blog project. After you have created your project directory, create virtual environment using the following commands:
- Windows
  ``` bash
  c:\>python -m venv c:\path\to\myenv
  ```
- Linux//MacOs
  ``` bash
  python3 -m venv /path/to/new/virtual/environment
  ```

Activate the virtual environment:
  - Windows cmd
  ``` bash
  C:\> <venv>\Scripts\activate.bat
  ```

  - Windows powershell
  ``` powershell
  <venv>\Scripts\Activate.ps1
  ```

  - Linux//MacOs
  ``` bash
  source <venv>/bin/activate
  ```

Now let's use `pip` to install required modules and packages that we will be using in this project.
``` bash
pip install flask markdown
```

## Creating the flask app

First, create a new Flask app, by creating a file in root of the project directory called `main.py`:

``` python
from flask import Flask, render_template
import markdown

app = Flask(__name__)
```

Define a route for the home page:
``` python
@app.route('/')
def home():
    return render_template('index.html')
```

Define a route to handle requests for individual blog posts:

``` python
@app.route('/posts/<path:path>')
def post(path):
    with open(f'posts/{path}.md', 'r') as file:
        content = file.read()
        html = markdown.markdown(content)
        return render_template('post.html', content=html)
```

Create templates for the home page and individual blog posts, we can do this by creating a new directory in root of project called `templates`. And then further create the two following `html` files:

- `index.html`:

``` html
<!DOCTYPE html>
<html>
<head>
    <title>My Blog</title>
</head>
<body>
    <h1>My Blog</h1>
    {% for post in posts %}
    <h2><a href="/posts/{{ post }}">{{ post }}</a></h2>
    {% endfor %}
</body>
</html>
```

- `post.html`:

``` html
<!DOCTYPE html>
<html>
<head>
    <title>{{ title }}</title>
</head>
<body>
    <h1>{{ title }}</h1>
    <div>{{ content|safe }}</div>
</body>
</html>
```

Modify the home route to display a list of blog post titles:

``` python
@app.route('/')
def home():
    posts = []
    for file in os.listdir('posts'):
        if file.endswith('.md'):
            title = file[:-3]
            posts.append(title)
    return render_template('index.html', posts=posts)
```

## Adding markdown posts

Now before running the app, let's add few posts.
Create a directory called `posts` and add some Markdown files with blog post content.
Let's add a `hello.md`:

``` markdown
# Hello

This is my first blog post
### Heading level 3
#### Heading level 4
##### Heading level 5
###### Heading level 6

I just love **bold text**.

```

Now, let's run the app, type the following command:

``` bash
python main.py
```

And you should see the following output in the termainal:

``` bash
 python main.py                                                                                                                * Serving Flask app 'main'                                                                                                     * Debug mode: on                                                                                                              WARNING: This is a development server. Do not use it in a production deployment. Use a production WSGI server instead.         
 * Running on http://127.0.0.1:5000
Press CTRL+C to quit
 * Restarting with stat
 * Debugger is active!
```

Here is how it would look, I have 2 blog posts and have some gifs in my blog posts. Navigate to `127.0.0.0:5000` in a browser window:

![Home Page of our blog](/2023/images/day48-1.png)

If we click on the `hello` blog post:

![Hello blog post](/2023/images/day48-2.png)

See you tomorrow in [Day 49](day49.md).

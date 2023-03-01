# Day 47 - Let's build an App in Python

Let's create a simple blog app with the help of [Flask](https://flask.palletsprojects.com/en/2.2.x/) that supports posts in [markdown.](https://www.markdownguide.org/basic-syntax/)

## Initiating virtual env and installing packages

Let's create a directory for our blog project. After you have created your project directory, create virtual environment using the following commands:
- Windows
  ``` bash
  ```
- Linux//MacOs
  ``` bash
  ```
  
Now let's use `pip` to install required modules and packages that we will be using in this project.
``` bash
pip install flask markdown
```

## Creating the flask app

First, create a new Flask app:

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

Create a directory called posts and add some Markdown files with blog post content.

Define a route to handle requests for individual blog posts:

``` python
@app.route('/posts/<path:path>')
def post(path):
    with open(f'posts/{path}.md', 'r') as file:
        content = file.read()
        html = markdown.markdown(content)
        return render_template('post.html', content=html)
```

Create templates for the home page and individual blog posts:
- `home.html`:

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

  

# Test Builder

A handy tool that allows you to build tests.

# Tutorial


Create a new file called `example.yml`. 

You start by naming the test using a `name:` element.

```YAML
name: Oregon Trail
```

Add a description with the `description:` element.

```YAML
description: A quiz on the Oregon Trail
```

Then you simply append a question. Notice how `is_answer: true` denotes which answer is correct.

```YAML
- question: Around how long is the Oregon Trail?
  type: multiple-choice
  answers:
  - answer: 100 miles
  - answer: 500 miles
  - answer: 1,000 miles
  - answer: 2,000 miles
    is_answer: true
  - answer: 4,000 miles
```

Simply continue to append questions and your test will grow. Here is the complete `example.yml`.

```YAML
name: Oregon Trail
description: A 10 question quiz on the Oregon Trail
questions:
- question: 'True or False: The main danger to pioneers on the trail was Native Americans.'
  type: true-false
  answers:
  - answer: "True"
  - answer: "False"
    is_answer: true
- question: Around how long is the Oregon Trail?
  type: multiple-choice
  answers:
  - answer: 100 miles
  - answer: 500 miles
  - answer: 1,000 miles
  - answer: 2,000 miles
    is_answer: true
  - answer: 4,000 miles
```

Now feed the file to TestBuilder.

```bash
$ ./tb -file example.yml
```

![](tb.gif)

# Developers

```bash
$ git clone https://github.com/jaredfolkins/tb
$ cd tb
$ go run *.go -example
$ go run *.go -file example.yml
```
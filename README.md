# Exercise #3: Choose your own adventure


## Details

[Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) is (was?) a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

- Turn to page 44 to go up the ladder.
- Turn to page 87 to venture down the dark passage.

The goal of this exercise is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

Stories will be provided via a JSON file with the following format:

```json
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```

*See [stories.json](https://github.com/Paul3435/CYOAdventure/blob/main/stories.json) for a real example of a JSON story. I find that seeing the real JSON file really helps answer any confusion or questions about the JSON format.*

Requirements:

1. Use the `html/template` package to create your HTML pages.
2. Create an `http.Handler` to handle the web requests instead of a handler function.
3. Use the `encoding/json` package to decode the JSON file.

A few things worth noting:

- Stories could be cyclical if a user chooses options that keep leading to the same place. This isn't likely to cause issues, but keep it in mind.
- For simplicity, all stories will have a story arc named "intro" that is where the story starts. That is, every JSON file will have a key with the value `intro` and this is where your story should start.
- Matt Holt's JSON-to-Go is a really handy tool when working with JSON in Go! Check it out - <https://mholt.github.io/json-to-go/>
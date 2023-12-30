## VisioLog: Searchable Screen History

VisioLog is a free and open-source software designed to help users maintain a searchable history of their screen activity. With a focus on simplicity and functionality, this tool automates the process of taking periodic screenshots and performs Optical Character Recognition (OCR) to extract text from these images. The goal is to enable users to easily search and find specific information from their screen history by entering keywords.  

> *Only **GNU/Linux** and **X11** are supported for now.*

### Key Features

**Automated Screen Capture:**
ScreenScribe captures screenshots at scheduled intervals, creating a chronological record of your screen activity.

**OCR Integration:**
The software utilizes OCR to convert text within the captured screenshots into searchable content.

**Basic Indexing:**
ScreenScribe employs basic indexing to organize and categorize the extracted text for easier retrieval.

**Search Functionality:**
Users can search for specific information within their screen history by entering keywords, allowing for quick access to relevant screenshots.

**Local Storage and Privacy:**
All data, including screenshots and extracted text, is stored locally to prioritize user privacy. No data is sent to external servers.

**Single Binary:**
Single binary with zero dependencies.

### Usage

Start capturing from your terminal:
```console
[~]$ ./visiolog capture
```

Access you screen history with built-in GUI
```console
[~]$ ./visiolog gui
```
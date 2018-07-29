# fo
File Organizer in Go

## Install
Install via go command:

```bash
$ go get -u github.com/wei0831/fo
```

## Examples

### folderin
#### Move Something_SXXEXX into their individual folder
```bash
$ folderin "D:\Video" -w
```
<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── Something_S01E01.mp4
├── Something_S01E01.jpg 
├── Something_S01E02.avi
            </code></pre></td>
            <td><pre><code>.
├── Something_S01E01
    ├── Something_S01E01.mp4
    ├── Something_S01E01.jpg 
├── Something_S01E02
    ├── Something_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

### folderout
#### Move files in the folder out
```bash
$ folderout "D:\Video" -w
```
<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── Something_S01E01
    ├── Something_S01E01.mp4
    ├── Something_S01E01.jpg 
├── Something_S01E02
    ├── Something_S01E02.avi
            </code></pre></td>
            <td><pre><code>.
├── Something_S01E01.mp4
├── Something_S01E01.jpg 
├── Something_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

### replacename
#### Remove [Bad] in file name only
```bash
$ fo replacename "\[Bad\]" "" -d "D:\Video" -w
```
<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [Bad]Something_Folder[Bad] 
├── [Bad]Something_S01E01[Bad].mp4
├── [Bad]Something_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── [Bad]Something_Folder[Bad]  
├── Something_S01E01.mp4
├── Something_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

#### Replace [Bad] in folder name obly
```bash
$ fo replacename "\[Bad\]" "" -d "D:\Video" -m1 -w
```
<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [Bad]Something_Folder[Bad]   
├── [Bad]Something_S01E01[Bad].mp4
├── [Bad]Something_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── Something_Folder
├── [Bad]Something_S01E01[Bad].mp4
├── [Bad]Something_S01E02[Bad].avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

#### Remove [Bad] in both folder name and file name
```bash
$ replacename "\[Bad\]" "" -d "D:\Video" -m2 -w
```
<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [Bad]Something_Folder[Bad]   
├── [Bad]Something_S01E01[Bad].mp4
├── [Bad]Something_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── Something_Folder
├── Something_S01E01.mp4
├── Something_S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

#### Change the name using regex group
```bash
$ replacename "(.*)(Something)(.*)(S[0-9]+E[0-9]+)(.*)(\.(mp4|avi))" "$2-$4$6" -d "D:\Something" -w
```
<table>
    <thead>
        <tr><th>Before</th><th>After</th></tr>
    </thead>
    <tbody>
        <tr>
            <td><pre><code>.
├── [20240202]Something_S01E01[Bad].mp4
├── [20240207]Something_S01E02[Bad].avi
            </code></pre></td>
            <td><pre><code>.
├── Something-S01E01.mp4
├── Something-S01E02.avi
            </code></pre></td>
        </tr>
    </tbody>
</table>

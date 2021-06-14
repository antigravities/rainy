import { Component } from "react";
import Dropzone from "react-dropzone";
import { Container } from "react-bootstrap";

class File {
    constructor(fx){
        this.fx = fx;
        this.status = "pre_upload";
        this.result = "";
    }

    async upload(){
        this.status = "uploading";

        const formData = new FormData();
        formData.append("file", this.fx);

        let req = await fetch("/upload", {
            method: 'POST',
            body: formData
        });

        this.result = await req.text();
        this.status = (! req.ok) ? "error" : "done";
    }
}

class Uploader extends Component {
    constructor(props){
        super(props);

        this.state = {};
        this.state.files = [];
    }

    async onDrop(files){
        files = files.map(i => new File(i));

        console.log(files);

        this.setState({ files: this.state.files.concat(files) });

        for( let file of files ) {
            await file.upload();

            this.setState({ files: this.state.files });
        }
    }

    parseStatus(file){
        switch(file.status){
            case "pre_upload":
                return <span>Preparing to upload</span>;
            case "uploading":
                return <span>Uploading</span>;
            case "done":
                return <a href={file.result} target="_blank" rel="noreferrer">{file.result}</a>;
            case "error":
                return <span>{file.result}</span>;
            default:
                return <span>{file.status}</span>;
        }
    }

    render(){
        return (
            <div>
                <Dropzone onDrop={files => this.onDrop(files)}>
                    {({getRootProps, getInputProps}) => (
                        <Container {...getRootProps({ className: 'dropzone' })}>
                            <input {...getInputProps()} />

                            Drop a file or click to upload...
                        </Container>
                    )}
                </Dropzone>

                <br></br>

                <h3>Uploads</h3>

                {this.state.files.length > 0 ? this.state.files.map(i => (<li key={i.fx.name}>{i.fx.name} - {this.parseStatus(i)}</li>)) : <i>There's nothing here!</i> }
            </div>
        )
    }
}

export default Uploader;
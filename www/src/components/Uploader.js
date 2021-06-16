import { Component } from "react";
import Dropzone from "react-dropzone";
import { Container } from "react-bootstrap";

class File {
    constructor(fx){
        this.fx = fx;
        this.status = "pre_upload";
        this.result = "";
    }

    async upload(password){
        this.status = "uploading";

        const formData = new FormData();
        formData.append("file", this.fx);

        // We pass the password as a GET variable here because if we want to parse a multipart request in Go,
        // we need to also hold on to the files that are part of that request. In theory this could fill up
        // the drive quickly or create situations where contraband is stored before password verification.
        let req = await fetch("/upload?password=" + encodeURIComponent(password), {
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

        this.setState({ files: this.state.files.concat(files) });

        for( let file of files ) {
            await file.upload(this.props.password);

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

    getShareX(){
        return btoa(`{
    "Version": "12.4.1",
    "DestinationType": "ImageUploader, TextUploader, FileUploader",
    "RequestMethod": "POST",
    "RequestURL": "${window.location.protocol}//${window.location.hostname}${window.location.port !== "" ? `:${window.location.port}` : ""}/upload${this.props.password ? `?password=${this.props.password}` : ""}",
    "Body": "MultipartFormData",
    "FileFormName": "file",
    "URL": "$response$"
}`);
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

                <br /><br />

                <h3>Uploads</h3>

                {this.state.files.length > 0 ? this.state.files.map(i => (<li key={i.fx.name}>{i.fx.name} - {this.parseStatus(i)}</li>)) : <i>There's nothing here!</i> }

                <br /><br />

                <a href={`data:text/plain;base64,${this.getShareX()}`} download={window.location.hostname + ".sxcu"}>Upload with ShareX</a>
            </div>
        )
    }
}

export default Uploader;
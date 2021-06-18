import { Component } from "react";
import Dropzone from "react-dropzone";
import { Container } from "react-bootstrap";
import UploadFile from "../lib/UploadFile";

class FileUploader extends Component {
    constructor(props){
        super(props);

        this.state = {};
        this.state.files = [];
    }

    async onDrop(files){
        files = files.map(i => new UploadFile(i));

        this.props.onFilesQueued(files);
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
            </div>
        )
    }
}

export default FileUploader;
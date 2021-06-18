import { Component } from "react";

class UploadedFiles extends Component {
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

    render() {
        return (
            <>
                <h3>Uploads</h3>

                {this.props.files.length > 0 ? this.props.files.map(i => (<li key={i.fx.name + "@" + i.time}>{i.fx.name} - {this.parseStatus(i)}</li>)) : "There's nothing here!"}
            </>
        )
    }
}

export default UploadedFiles;
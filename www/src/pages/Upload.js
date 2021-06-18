import { Component } from "react";
import { Container } from 'react-bootstrap';
import UploadTabs from "../components/UploadTabs";
import UploadedFiles from "../components/UploadedFiles";
import PasswordEntryForm from "../components/PasswordEntryForm";

class PageUpload extends Component {
    constructor(props){
        super(props);

        this.state = {};

        // a bit of a hack, but strings coerce to boolean true
        this.state.uploadPassword = props.meta.hasUploadPassword;
        this.state.files = [];
    }

    updateUploadPassword(password){
        this.setState({ uploadPassword: password });
    }

    async onFilesQueued(files){
        this.setState({ files: this.state.files.concat(files) });

        for( let file of files ) {
            await file.upload(this.state.uploadPassword);
            this.setState({ files: this.state.files });
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

    render() {
        return (
            <Container>
                <h1>{this.props.meta.tagline}</h1>

                Any files up to {this.props.meta.maxUploadSizeMB} MB are accepted{this.props.meta.extensionBlacklist.length > 0 ? <span> (except {"." + this.props.meta.extensionBlacklist.split(",").join(", .")} files)</span> : ""}.

                <br /><br />

                {this.props.meta.hasUploadPassword && ! this.state.uploadPassword ?
                    <PasswordEntryForm meta={this.props.meta} onPasswordEntered={pw => this.updateUploadPassword(pw)} /> :
                    <><UploadTabs onFilesQueued={f => this.onFilesQueued(f)} /><hr /><UploadedFiles files={this.state.files} /></>
                }
            </Container>
        )
    }
}

export default PageUpload;
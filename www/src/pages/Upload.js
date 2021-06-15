import { Component } from "react";
import { Container } from 'react-bootstrap';
import Uploader from "../components/Uploader";
import PasswordEntryForm from "../components/PasswordEntryForm";

class PageUpload extends Component {
    constructor(props){
        super(props);

        this.state = {};
        this.state.hasUploadPassword = props.meta.hasUploadPassword;
    }

    updateUploadPassword(password){
        this.setState({ hasUploadPassword: password });
    }

    render() {
        return (
            <Container>
                <h1>{this.props.meta.tagline}</h1>

                Any files up to {this.props.meta.maxUploadSizeMB} MB are accepted{this.props.meta.extensionBlacklist.length > 0 ? <span> (except {"." + this.props.meta.extensionBlacklist.split(",").join(", .")} files)</span> : ""}.

                <br /><br />

                {this.props.meta.hasUploadPassword && ! this.state.hasUploadPassword ? <PasswordEntryForm meta={this.props.meta} onPasswordEntered={pw => this.updateUploadPassword(pw)} /> : <Uploader password={this.state.hasUploadPassword} />}

            </Container>
        )
    }
}

export default PageUpload;
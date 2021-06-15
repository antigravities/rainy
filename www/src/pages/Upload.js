import { Component } from "react";
import { Container, Form } from 'react-bootstrap';
import Header from '../components/Header';
import Footer from '../components/Footer';
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
            <>
                <Header title={this.props.meta.instanceName} />

                <Container>
                    <h1>Upload your freshest mee-mee, fam.</h1>

                    Any files up to {this.props.meta.maxUploadSizeMB} MB are accepted{this.props.meta.extensionBlacklist.length > 0 ? <span> (except {"." + this.props.meta.extensionBlacklist.split(",").join(", .")} files)</span> : ""}.

                    <br /><br />

                    {this.props.meta.hasUploadPassword && ! this.state.hasUploadPassword ? <PasswordEntryForm meta={this.props.meta} onPasswordEntered={pw => this.updateUploadPassword(pw)} /> : <Uploader password={this.state.hasUploadPassword} />}

                </Container>

                <Footer />
            </>
        )
    }
}

export default PageUpload;
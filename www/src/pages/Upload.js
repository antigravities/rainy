import { Component } from "react";
import { Container } from 'react-bootstrap';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Uploader from "../components/Uploader";

class PageUpload extends Component {
    render() {
        return (
            <>
                <Header title={this.props.meta.instanceName} />

                <Container>
                    <h1>Upload your freshest mee-mee, fam.</h1>

                    Any files up to {this.props.meta.maxUploadSizeMB} MB are accepted{this.props.meta.extensionBlacklist.length > 0 ? <span> (except {"." + this.props.meta.extensionBlacklist.split(",").join(", .")} files)</span> : ""}.

                    <br /><br />

                    <Uploader />

                </Container>

                <Footer />
            </>
        )
    }
}

export default PageUpload;
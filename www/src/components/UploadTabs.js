import { Component } from "react";
import { Tab, Tabs } from "react-bootstrap";
import FileUploader from "./FileUploader";
import TextUploader from "./TextUploader";

class UploadTabs extends Component {
    render(){
        return (
            <Tabs defaultActiveKey="file" id="upload-tabs">
                <Tab eventKey="file" title="File" className="uploaderTab">
                    <FileUploader onFilesQueued={this.props.onFilesQueued} />
                </Tab>
                <Tab eventKey="text" title="Text" className="uploaderTab">
                    <TextUploader onFilesQueued={this.props.onFilesQueued} />
                </Tab>
            </Tabs>
        )
    }
}

export default UploadTabs;
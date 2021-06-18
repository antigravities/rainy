import { Component } from "react";
import { Button } from "react-bootstrap";
import UploadFile from "../lib/UploadFile";

class TextUploader extends Component {
    constructor(props){
        super(props);

        this.state = {};
        this.state.value = "";
    }

    handleChange(e){
        this.setState({ value: e.target.value });
    }

    handleSubmit(e){
        e.preventDefault();
        if( this.state.value.length < 1 ) return;
        this.props.onFilesQueued([new UploadFile(new File([new Blob([this.state.value])], "file.txt"))]);
        this.setState({ value: "" });
    }

    render(){
        return (
            <form onSubmit={e => this.handleSubmit(e)}>
                <textarea className="uploadText" placeholder="Paste or type text here..." onChange={e => this.handleChange(e)} value={this.state.value}></textarea>

                <br /><br />

                <Button type="submit">Submit</Button>
            </form>
        )
    }
}

export default TextUploader;
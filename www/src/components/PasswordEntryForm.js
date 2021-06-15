import { Component } from "react";
import { Form, Button } from "react-bootstrap"

class PasswordEntryForm extends Component {
    async onSubmit(e){
        e.preventDefault();

        let elem = document.querySelector("#uploadPassword");

        let res = await fetch("/meta", {
            method: "POST",
            body: "password=" + window.encodeURIComponent(elem.value),
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            }
        });

        if( res.ok ){
            this.props.onPasswordEntered(elem.value);
        } else {
            elem.setAttribute("placeholder", await res.text())
            elem.value = "";
            elem.focus();
        }
    }

    render(){
        return (<Form onSubmit={e => this.onSubmit(e)}>
            <Form.Group>
                <Form.Label>{this.props.meta.instanceName} restricts who can upload to it with a password. Enter it below to gain access to the upload form.</Form.Label>
                <Form.Control type="password" placeholder="Upload password" id="uploadPassword"></Form.Control>

                <br />

                <Button variant="primary" type="submit">
                    Submit
                </Button>
            </Form.Group>
        </Form>);
    }
}

export default PasswordEntryForm;
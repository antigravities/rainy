import { Component } from "react";
import { Container } from "react-bootstrap";

class Footer extends Component {
    render(){
        return (
            <Container>
                <hr />

                &copy; 2021 <a href="https://cutie.cafe/" target="_blank" rel="noreferrer">Cutie Caf&eacute;.</a> All rights reserved.<br />
            </Container>
        )
    }
}

export default Footer;
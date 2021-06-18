import { Component } from "react";
import { Container } from "react-bootstrap";

class Footer extends Component {
    render(){
        return (
            <Container>
                <hr />

                <span className="text-muted">Powered by <a href="https://github.com/antigravities/rainy" target="_blank" rel="noreferrer">Rainy</a> &middot; <a href="/upload">No JS mode</a></span>
            </Container>
        )
    }
}

export default Footer;
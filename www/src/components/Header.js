import { Component } from "react";
import { Navbar, Nav } from "react-bootstrap";

class Header extends Component {
    render() {
        return (
            <Navbar bg="light" expand="lg">
                <Navbar.Brand href="/">
                    {this.props.title}
                </Navbar.Brand>
                
                <Navbar.Toggle aria-controls="nav" />

                <Navbar.Collapse id="nav">
                    <Nav className="me-auto">
                        <Nav.Link href="/">
                            Upload
                        </Nav.Link>
                    </Nav>
                </Navbar.Collapse>
            </Navbar>
        )
    }
}

export default Header;
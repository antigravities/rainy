import React, { Component } from 'react';
import './App.css';
import Upload from './pages/Upload';

class App extends Component {
  constructor(props){
    super(props);

    this.state = {};
    this.state.meta = { extensionBlacklist: "" };
    this.state.showModal = false;
  }

  async componentDidMount(){
    let meta = await fetch("/meta", {
      headers: {
        "Accept": "application/json"
      },
      cache: "no-cache"
    });

    this.setState({meta: await meta.json()});

    window.document.title = this.state.meta.instanceName;
  }

  async showModal(){
    this.setState({ showModal: true });
  }

  render() {
    return (
      <div>
        <Upload meta={this.state.meta}>
          
        </Upload>
      </div>
    )
  }
}

export default App;

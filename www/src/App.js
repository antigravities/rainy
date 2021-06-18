import React, { Component } from 'react';
import './App.css';
import Upload from './pages/Upload';
import Header from './components/Header';
import Footer from './components/Footer';

class App extends Component {
  constructor(props){
    super(props);

    this.state = {};
    this.state.meta = { extensionBlacklist: "" };
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

  render() {
    return (
      <div>
        <Header title={this.state.meta.instanceName} />
        <Upload meta={this.state.meta} />
        <Footer />
      </div>
    )
  }
}

export default App;

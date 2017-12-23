import React, { Component } from 'react';
import { BrowserRouter, Route, Link } from 'react-router-dom';
import ReactDOM from 'react-dom';

// import style
require('./style.scss');

const Home = () => (
  <div>
    <h2>Home</h2>
  </div>
)

const About = () => (
  <div>
    <h2>About</h2>
  </div>
)

const App = () => (
  <div className="wrapper">
    <nav>
      <Link to="/">Home</Link>
      <Link to="/about">About</Link>
    </nav>
    <div className="content">
      <Route exact path="/" component={Home} />
      <Route path="/about" component={About} />
    </div>
  </div>
)

class Container extends Component {

  render() {
    return (
      <BrowserRouter>
        <App />
      </BrowserRouter>
    )
  }

}

ReactDOM.render(<Container />, document.querySelector('#app'));

if (module.hot) {
  module.hot.accept();
}

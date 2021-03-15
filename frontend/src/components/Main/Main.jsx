import React from "react";
import { SearchBar } from "../SearchBar";
import { DownloadItem } from "../DownloadItem";
import "./Main.css";

class Main extends React.Component {
  constructor(props) {
    super(props);
    this.state = { url: "", loading: false };
    this.passValue = this.passValue.bind(this);
    this.passLoadingStatus = this.passLoadingStatus.bind(this);
  }

  passValue(value) {
    this.setState({ url: value });
  }

  passLoadingStatus(value) {
    this.setState({ loading: value });
  }

  render() {
    let item;
    if (this.state.url != null && this.state.url !== "") {
      item = <DownloadItem url={this.state.url} />;
    } else {
      item = "";
    }

    return (
      <div className="container">
        <br />
        <br />
        <br />
        <div className="row">
          <div className="col-md-6 offset-md-3">
            <SearchBar
              onClickHandler={this.passValue}
              onLoadingHandler={this.passLoadingStatus}
            />
          </div>
          <div className="col-md-6 offset-md-3">{item}</div>
        </div>
        <div
          id="overlay"
          style={{ display: this.state.loading ? "block" : "none" }}
        >
          <div className="spinner"></div>
          <br />
          Loading...
        </div>
      </div>
    );
  }
}

export { Main };

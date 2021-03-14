import React from "react";
import { SearchBar } from "../SearchBar";
import { DownloadItem } from "../DownloadItem";

class Main extends React.Component {
  constructor(props) {
    super(props);
    this.state = { url: "" };
    this.passValue = this.passValue.bind(this);
  }

  passValue(value) {
    this.setState({ url: value });
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
            <SearchBar onClickHandler={this.passValue} />
          </div>
          <div className="col-md-6 offset-md-3">{item}</div>
        </div>
      </div>
    );
  }
}

export { Main };

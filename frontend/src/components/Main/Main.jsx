import React from "react";
import { SearchBar } from "../SearchBar";
import { DownloadItem } from "../DownloadItem";

class Main extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="row">
          <div className="col-md-6 offset-md-3">
            <SearchBar />
          </div>
          <div className="col-md-6 offset-md-3">
            <DownloadItem />
          </div>
        </div>
      </div>
    );
  }
}

export { Main };

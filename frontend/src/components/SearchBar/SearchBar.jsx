import React from "react";

class SearchBar extends React.Component {
  render() {
    return (
      <div className="input-group mb-3">
        <input
          type="text"
          className="form-control"
          placeholder="搜尋關鍵字"
          aria-label="搜尋關鍵字"
          aria-describedby="basic-addon2"
        />
        <span className="input-group-text" id="basic-addon2">
          搜尋
        </span>
      </div>
    );
  }
}

export { SearchBar };

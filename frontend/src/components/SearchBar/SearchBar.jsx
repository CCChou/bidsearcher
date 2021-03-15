import React from "react";
import PropTypes from "prop-types";

class SearchBar extends React.Component {
  constructor(props) {
    super(props);
    this.state = { text: "" };
    this.onClickHandler = props.onClickHandler;
    this.onLoadingHandler = props.onLoadingHandler;
    this.getTextValue = this.getTextValue.bind(this);
  }

  static get propTypes() {
    return {
      onClickHandler: PropTypes.func,
      onLoadingHandler: PropTypes.func,
    };
  }

  getTextValue() {
    let searchBar = this;
    searchBar.onLoadingHandler(true);
    fetch(window.location.href + "search?keyword=" + this.state.text)
      .then((res) => res.json())
      .then(function (jsonData) {
        searchBar.onLoadingHandler(false);
        if (jsonData.status === "OK") {
          searchBar.onClickHandler(jsonData.url);
        }
      })
      .catch((error) => {
        searchBar.onLoadingHandler(false);
        console.error("Error: ", error);
      });
  }

  render() {
    return (
      <div className="input-group mb-3">
        <input
          type="text"
          className="form-control"
          placeholder="搜尋關鍵字"
          aria-label="搜尋關鍵字"
          aria-describedby="basic-addon2"
          onChange={(e) => this.setState({ text: e.target.value })}
        />
        <button
          className="input-group-text btn btn-primary"
          id="basic-addon2"
          onClick={this.getTextValue}
        >
          搜尋
        </button>
      </div>
    );
  }
}

export { SearchBar };

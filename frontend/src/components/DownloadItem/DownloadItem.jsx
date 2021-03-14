import React from "react";
import PropTypes from "prop-types";

class DownloadItem extends React.Component {
  static get propTypes() {
    return {
      url: PropTypes.string,
    };
  }

  render() {
    return (
      <a
        href={this.props.url}
        className="list-group-item list-group-item-action list-group-item-primary"
      >
        {this.props.url.split("/").pop()}
      </a>
    );
  }
}

export { DownloadItem };

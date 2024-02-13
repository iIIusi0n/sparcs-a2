import React from "react";
import PropTypes from "prop-types";

const MyComponents = ({ name, children }) => {
  return (
    <div>
      {name}
      children값은 {children}이다.
    </div>
  );
};

MyComponents.defaultProps = {
  name: "기본이름",
};

export default MyComponents;

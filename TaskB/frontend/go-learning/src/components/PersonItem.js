import React from 'react';

function PersonItem({ person }) {
  return (
    <li className="person-item">
      <span className="person-name">{person.name}</span>
      <span className="person-age">Age: {person.age}</span>
    </li>
  );
}

export default PersonItem;

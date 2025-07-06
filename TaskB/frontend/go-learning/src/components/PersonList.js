import React from 'react';
import PersonItem from './PersonItem';

function PersonList({ persons }) {
  if (!Array.isArray(persons)) return <p>No persons available.</p>;

  return (
    <ul className="person-list">
      {persons.map((person) => (
        <PersonItem key={person.id} person={person} />
      ))}
    </ul>
  );
}

export default PersonList;

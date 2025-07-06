import React, { useState, useEffect } from 'react';
import axios from 'axios';
import PersonList from '../components/PersonList'; // Adjust path if needed
import PaginationControls from '../components/PaginationControls';

function PersonListPage() {
  const [persons, setPersons] = useState([]);
  const [page, setPage] = useState(1);
  const [pageSize] = useState(5);
  const [hasNext, setHasNext] = useState(false);

  useEffect(() => {
    const fetchPersons = async () => {
      try {
        const { data } = await axios.get(`http://localhost:8080/persons?page=${page}&page_size=${pageSize}`);
        setPersons(data.data);
        setHasNext(data.has_next);
      } catch (err) {
        console.error('Error fetching persons:', err);
      }
    };

    fetchPersons();
  }, [page, pageSize]);

  return (
    <div className="app-container">
      <h1 className="header">Persons List</h1>
      <PersonList persons={persons} />
      <PaginationControls page={page} setPage={setPage} hasNext={hasNext} />
    </div>
  );
}

export default PersonListPage;

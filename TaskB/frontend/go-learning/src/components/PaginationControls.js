import React from 'react';

function PaginationControls({ page, setPage, hasNext }) {
  return (
    <div className="pagination-controls">
      <button onClick={() => setPage(page - 1)} disabled={page === 1}>
        ⬅ Prev
      </button>
      <span className="page-indicator">Page {page}</span>
      <button onClick={() => setPage(page + 1)} disabled={!hasNext}>
        Next ➡
      </button>
    </div>
  );
}

export default PaginationControls;

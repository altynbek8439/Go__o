import { render, screen } from '@testing-library/react';
import App from './App';

test('renders learn react link', () => {
  render(<App />);
  const linkElement = screen.getByText(/learn react/i);
  expect(linkElement).toBeInTheDocument();
});
// import { render, screen } from '@testing-library/react';
// import App from './App';

// test('renders app title', () => {
//     render(<App />);
//     const titleElement = screen.getByText(/Букмекерский сайт/i);
//     expect(titleElement).toBeInTheDocument();
// });
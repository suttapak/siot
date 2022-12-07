import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { Container } from './components';
import { NotAuth, ProviderAuth, RequireAuth } from './hooks';
import { ProviderToast } from './hooks/useToast';
import { BoxDashBoard, SignIn, SignUp } from './pages';
import { HomePage } from './pages/Home.Page';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import BoxPage from './pages/BoxPage';
import BoxSecret from './pages/BoxSecretPage';
import BoxMemberPage from './pages/BoxMemberPage';
import { ProviderSokcetIO } from './hooks/useSocketIO';

const queryClient = new QueryClient();

function App() {
  return (
    <React.Fragment>
      <ProviderAuth>
        <ProviderToast>
          <QueryClientProvider client={queryClient}>
            <ProviderSokcetIO>
              <Container>
                <Routes>
                  <Route
                    path='/'
                    element={
                      <RequireAuth>
                        <HomePage />
                      </RequireAuth>
                    }
                  />
                  <Route
                    path='/boxes/:boxId'
                    element={
                      <RequireAuth>
                        <BoxPage />
                      </RequireAuth>
                    }
                  />
                  <Route
                    path='/boxes/:boxId/dashboard'
                    element={
                      <RequireAuth>
                        <BoxDashBoard />
                      </RequireAuth>
                    }
                  />
                  <Route
                    path='/boxes/:boxId/secret'
                    element={
                      <RequireAuth>
                        <BoxSecret />
                      </RequireAuth>
                    }
                  />
                  <Route
                    path='/boxes/:boxId/members'
                    element={
                      <RequireAuth>
                        <BoxMemberPage />
                      </RequireAuth>
                    }
                  />
                  <Route
                    path='signin'
                    element={
                      <NotAuth>
                        <SignIn />
                      </NotAuth>
                    }
                  />
                  <Route
                    path='signup'
                    element={
                      <NotAuth>
                        <SignUp />
                      </NotAuth>
                    }
                  />
                </Routes>
              </Container>
            </ProviderSokcetIO>
          </QueryClientProvider>
        </ProviderToast>
      </ProviderAuth>
    </React.Fragment>
  );
}

export default App;

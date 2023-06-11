import * as React from 'react';
import AuthLayout from 'layouts/Auth';

type AuthProps = { children: JSX.Element };

const Auth = ({ children }: AuthProps): JSX.Element => <AuthLayout>{children}</AuthLayout>;

export default Auth;

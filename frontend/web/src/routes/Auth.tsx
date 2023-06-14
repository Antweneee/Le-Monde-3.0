import * as React from 'react';
import AuthLayout from 'layouts/Auth';
import { Outlet } from 'react-router-dom';

const Auth = (): JSX.Element => (
	<AuthLayout>
		<Outlet />
	</AuthLayout>
);

export default Auth;

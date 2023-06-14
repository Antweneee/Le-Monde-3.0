import * as React from 'react';
import PrivateLayout from 'layouts/Private';
import { Outlet } from 'react-router-dom';

const Private = (): JSX.Element => (
	<PrivateLayout>
		<Outlet />
	</PrivateLayout>
);

export default Private;

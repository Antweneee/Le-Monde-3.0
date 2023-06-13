import * as React from 'react';
import PrivateLayout from 'layouts/Private';

type PrivateProps = { children: JSX.Element };

const Private = ({ children }: PrivateProps): JSX.Element => <PrivateLayout>{children}</PrivateLayout>;

export default Private;

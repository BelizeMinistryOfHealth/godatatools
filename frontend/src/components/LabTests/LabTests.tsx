import React from 'react';
import { LabTest } from '../../models/labTest';
import {
  Table,
  TableBody,
  TableCell,
  TableHeader,
  TableRow,
  Text,
} from 'grommet';

const rows = (test: LabTest) => {
  return (
    <TableRow key={test.id}>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true}>
          {test.person.firstName}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true}>
          {test.person.lastName}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true} wordBreak={'break-all'}>
          {test.sampleIdentifier}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true}>
          {test.dateSampleTaken}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true}>
          {test.dateTesting}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true}>
          {test.dateOfResult}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true} wordBreak={'break-all'}>
          {test.testType}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true} wordBreak={'break-all'}>
          {test.sampleType}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true} wordBreak={'break-all'}>
          {test.status}
        </Text>
      </TableCell>
      <TableCell wrap={true}>
        <Text size={'small'} truncate={true} wordBreak={'break-all'}>
          {test.result}
        </Text>
      </TableCell>
    </TableRow>
  );
};

interface LabTestTableProps {
  labTests: LabTest[];
}
const LabTestTable = (props: LabTestTableProps) => {
  const { labTests } = props;
  return (
    <Table>
      <TableHeader>
        <TableRow>
          <TableCell>
            <Text size={'small'}>First Name</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Last Name</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Sample Identifier</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Sample Date</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Test Date</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Result Date</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Test Type</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Sample Type</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Status</Text>
          </TableCell>
          <TableCell>
            <Text size={'small'}>Result</Text>
          </TableCell>
        </TableRow>
      </TableHeader>
      <TableBody>{labTests.map((l) => rows(l))}</TableBody>
    </Table>
  );
};

export interface LabTestsProps {
  labTests: LabTest[];
}
const LabTests = (props: LabTestsProps) => {
  const { labTests } = props;
  return <LabTestTable labTests={labTests} />;
};

export default LabTests;

export interface LabTest {
  id: string;
  person: Person;
  dateSampleTaken: Date;
  dateSampleDelivered: Date;
  dateTesting: Date;
  dateOfResult: Date;
  sampleType: string;
  testType: string;
  result: string;
  status: string;
  sampleIdentifier: string;
  labFacility: LabFacility;
}

export interface Person {
  firstName: string;
  lastName: string;
  gender: string;
  dob: Date;
  age: number;
}

export interface LabFacility {
  id: string;
  name: string;
}

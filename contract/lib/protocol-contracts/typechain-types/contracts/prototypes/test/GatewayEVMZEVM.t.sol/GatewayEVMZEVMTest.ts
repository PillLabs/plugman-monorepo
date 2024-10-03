/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "../../../../common";

export declare namespace StdInvariant {
  export type FuzzSelectorStruct = {
    addr: PromiseOrValue<string>;
    selectors: PromiseOrValue<BytesLike>[];
  };

  export type FuzzSelectorStructOutput = [string, string[]] & {
    addr: string;
    selectors: string[];
  };

  export type FuzzArtifactSelectorStruct = {
    artifact: PromiseOrValue<string>;
    selectors: PromiseOrValue<BytesLike>[];
  };

  export type FuzzArtifactSelectorStructOutput = [string, string[]] & {
    artifact: string;
    selectors: string[];
  };

  export type FuzzInterfaceStruct = {
    addr: PromiseOrValue<string>;
    artifacts: PromiseOrValue<string>[];
  };

  export type FuzzInterfaceStructOutput = [string, string[]] & {
    addr: string;
    artifacts: string[];
  };
}

export interface GatewayEVMZEVMTestInterface extends utils.Interface {
  functions: {
    "IS_TEST()": FunctionFragment;
    "excludeArtifacts()": FunctionFragment;
    "excludeContracts()": FunctionFragment;
    "excludeSelectors()": FunctionFragment;
    "excludeSenders()": FunctionFragment;
    "failed()": FunctionFragment;
    "setUp()": FunctionFragment;
    "targetArtifactSelectors()": FunctionFragment;
    "targetArtifacts()": FunctionFragment;
    "targetContracts()": FunctionFragment;
    "targetInterfaces()": FunctionFragment;
    "targetSelectors()": FunctionFragment;
    "targetSenders()": FunctionFragment;
    "testCallReceiverEVMFromSenderZEVM()": FunctionFragment;
    "testCallReceiverEVMFromZEVM()": FunctionFragment;
    "testWithdrawAndCallReceiverEVMFromSenderZEVM()": FunctionFragment;
    "testWithdrawAndCallReceiverEVMFromZEVM()": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "IS_TEST"
      | "excludeArtifacts"
      | "excludeContracts"
      | "excludeSelectors"
      | "excludeSenders"
      | "failed"
      | "setUp"
      | "targetArtifactSelectors"
      | "targetArtifacts"
      | "targetContracts"
      | "targetInterfaces"
      | "targetSelectors"
      | "targetSenders"
      | "testCallReceiverEVMFromSenderZEVM"
      | "testCallReceiverEVMFromZEVM"
      | "testWithdrawAndCallReceiverEVMFromSenderZEVM"
      | "testWithdrawAndCallReceiverEVMFromZEVM"
  ): FunctionFragment;

  encodeFunctionData(functionFragment: "IS_TEST", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "excludeArtifacts",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "excludeContracts",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "excludeSelectors",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "excludeSenders",
    values?: undefined
  ): string;
  encodeFunctionData(functionFragment: "failed", values?: undefined): string;
  encodeFunctionData(functionFragment: "setUp", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "targetArtifactSelectors",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "targetArtifacts",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "targetContracts",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "targetInterfaces",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "targetSelectors",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "targetSenders",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "testCallReceiverEVMFromSenderZEVM",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "testCallReceiverEVMFromZEVM",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "testWithdrawAndCallReceiverEVMFromSenderZEVM",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "testWithdrawAndCallReceiverEVMFromZEVM",
    values?: undefined
  ): string;

  decodeFunctionResult(functionFragment: "IS_TEST", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "excludeArtifacts",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "excludeContracts",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "excludeSelectors",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "excludeSenders",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "failed", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "setUp", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "targetArtifactSelectors",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "targetArtifacts",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "targetContracts",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "targetInterfaces",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "targetSelectors",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "targetSenders",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "testCallReceiverEVMFromSenderZEVM",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "testCallReceiverEVMFromZEVM",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "testWithdrawAndCallReceiverEVMFromSenderZEVM",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "testWithdrawAndCallReceiverEVMFromZEVM",
    data: BytesLike
  ): Result;

  events: {
    "Call(address,bytes,bytes)": EventFragment;
    "Call(address,address,bytes)": EventFragment;
    "Deposit(address,address,uint256,address,bytes)": EventFragment;
    "Executed(address,uint256,bytes)": EventFragment;
    "ExecutedWithERC20(address,address,uint256,bytes)": EventFragment;
    "ReceivedERC20(address,uint256,address,address)": EventFragment;
    "ReceivedNoParams(address)": EventFragment;
    "ReceivedNonPayable(address,string[],uint256[],bool)": EventFragment;
    "ReceivedPayable(address,uint256,string,uint256,bool)": EventFragment;
    "Withdrawal(address,address,bytes,uint256,uint256,uint256,bytes)": EventFragment;
    "log(string)": EventFragment;
    "log_address(address)": EventFragment;
    "log_array(uint256[])": EventFragment;
    "log_array(int256[])": EventFragment;
    "log_array(address[])": EventFragment;
    "log_bytes(bytes)": EventFragment;
    "log_bytes32(bytes32)": EventFragment;
    "log_int(int256)": EventFragment;
    "log_named_address(string,address)": EventFragment;
    "log_named_array(string,uint256[])": EventFragment;
    "log_named_array(string,int256[])": EventFragment;
    "log_named_array(string,address[])": EventFragment;
    "log_named_bytes(string,bytes)": EventFragment;
    "log_named_bytes32(string,bytes32)": EventFragment;
    "log_named_decimal_int(string,int256,uint256)": EventFragment;
    "log_named_decimal_uint(string,uint256,uint256)": EventFragment;
    "log_named_int(string,int256)": EventFragment;
    "log_named_string(string,string)": EventFragment;
    "log_named_uint(string,uint256)": EventFragment;
    "log_string(string)": EventFragment;
    "log_uint(uint256)": EventFragment;
    "logs(bytes)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "Call(address,bytes,bytes)"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "Call(address,address,bytes)"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Deposit"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Executed"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ExecutedWithERC20"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ReceivedERC20"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ReceivedNoParams"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ReceivedNonPayable"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "ReceivedPayable"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Reverted"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "RevertedWithERC20"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "Withdrawal"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_address"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_array(uint256[])"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_array(int256[])"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_array(address[])"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_bytes"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_bytes32"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_int"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_address"): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "log_named_array(string,uint256[])"
  ): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "log_named_array(string,int256[])"
  ): EventFragment;
  getEvent(
    nameOrSignatureOrTopic: "log_named_array(string,address[])"
  ): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_bytes"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_bytes32"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_decimal_int"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_decimal_uint"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_int"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_string"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_named_uint"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_string"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "log_uint"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "logs"): EventFragment;
}

export interface Call_address_bytes_bytes_EventObject {
  sender: string;
  receiver: string;
  message: string;
}
export type Call_address_bytes_bytes_Event = TypedEvent<
  [string, string, string],
  Call_address_bytes_bytes_EventObject
>;

export type Call_address_bytes_bytes_EventFilter =
  TypedEventFilter<Call_address_bytes_bytes_Event>;

export interface Call_address_address_bytes_EventObject {
  sender: string;
  receiver: string;
  payload: string;
}
export type Call_address_address_bytes_Event = TypedEvent<
  [string, string, string],
  Call_address_address_bytes_EventObject
>;

export type Call_address_address_bytes_EventFilter =
  TypedEventFilter<Call_address_address_bytes_Event>;

export interface DepositEventObject {
  sender: string;
  receiver: string;
  amount: BigNumber;
  asset: string;
  payload: string;
}
export type DepositEvent = TypedEvent<
  [string, string, BigNumber, string, string],
  DepositEventObject
>;

export type DepositEventFilter = TypedEventFilter<DepositEvent>;

export interface ExecutedEventObject {
  destination: string;
  value: BigNumber;
  data: string;
}
export type ExecutedEvent = TypedEvent<
  [string, BigNumber, string],
  ExecutedEventObject
>;

export type ExecutedEventFilter = TypedEventFilter<ExecutedEvent>;

export interface ExecutedWithERC20EventObject {
  token: string;
  to: string;
  amount: BigNumber;
  data: string;
}
export type ExecutedWithERC20Event = TypedEvent<
  [string, string, BigNumber, string],
  ExecutedWithERC20EventObject
>;

export type ExecutedWithERC20EventFilter =
  TypedEventFilter<ExecutedWithERC20Event>;

export interface ReceivedERC20EventObject {
  sender: string;
  amount: BigNumber;
  token: string;
  destination: string;
}
export type ReceivedERC20Event = TypedEvent<
  [string, BigNumber, string, string],
  ReceivedERC20EventObject
>;

export type ReceivedERC20EventFilter = TypedEventFilter<ReceivedERC20Event>;

export interface ReceivedNoParamsEventObject {
  sender: string;
}
export type ReceivedNoParamsEvent = TypedEvent<
  [string],
  ReceivedNoParamsEventObject
>;

export type ReceivedNoParamsEventFilter =
  TypedEventFilter<ReceivedNoParamsEvent>;

export interface ReceivedNonPayableEventObject {
  sender: string;
  strs: string[];
  nums: BigNumber[];
  flag: boolean;
}
export type ReceivedNonPayableEvent = TypedEvent<
  [string, string[], BigNumber[], boolean],
  ReceivedNonPayableEventObject
>;

export type ReceivedNonPayableEventFilter =
  TypedEventFilter<ReceivedNonPayableEvent>;

export interface ReceivedPayableEventObject {
  sender: string;
  value: BigNumber;
  str: string;
  num: BigNumber;
  flag: boolean;
}
export type ReceivedPayableEvent = TypedEvent<
  [string, BigNumber, string, BigNumber, boolean],
  ReceivedPayableEventObject
>;

export type ReceivedPayableEventFilter = TypedEventFilter<ReceivedPayableEvent>;

export interface RevertedEventObject {
  destination: string;
  value: BigNumber;
  data: string;
}
export type RevertedEvent = TypedEvent<
  [string, BigNumber, string],
  RevertedEventObject
>;

export type RevertedEventFilter = TypedEventFilter<RevertedEvent>;

export interface RevertedWithERC20EventObject {
  token: string;
  to: string;
  amount: BigNumber;
  data: string;
}
export type RevertedWithERC20Event = TypedEvent<
  [string, string, BigNumber, string],
  RevertedWithERC20EventObject
>;

export type RevertedWithERC20EventFilter =
  TypedEventFilter<RevertedWithERC20Event>;

export interface WithdrawalEventObject {
  from: string;
  zrc20: string;
  to: string;
  value: BigNumber;
  gasfee: BigNumber;
  protocolFlatFee: BigNumber;
  message: string;
}
export type WithdrawalEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber, BigNumber, string],
  WithdrawalEventObject
>;

export type WithdrawalEventFilter = TypedEventFilter<WithdrawalEvent>;

export interface logEventObject {
  arg0: string;
}
export type logEvent = TypedEvent<[string], logEventObject>;

export type logEventFilter = TypedEventFilter<logEvent>;

export interface log_addressEventObject {
  arg0: string;
}
export type log_addressEvent = TypedEvent<[string], log_addressEventObject>;

export type log_addressEventFilter = TypedEventFilter<log_addressEvent>;

export interface log_array_uint256_array_EventObject {
  val: BigNumber[];
}
export type log_array_uint256_array_Event = TypedEvent<
  [BigNumber[]],
  log_array_uint256_array_EventObject
>;

export type log_array_uint256_array_EventFilter =
  TypedEventFilter<log_array_uint256_array_Event>;

export interface log_array_int256_array_EventObject {
  val: BigNumber[];
}
export type log_array_int256_array_Event = TypedEvent<
  [BigNumber[]],
  log_array_int256_array_EventObject
>;

export type log_array_int256_array_EventFilter =
  TypedEventFilter<log_array_int256_array_Event>;

export interface log_array_address_array_EventObject {
  val: string[];
}
export type log_array_address_array_Event = TypedEvent<
  [string[]],
  log_array_address_array_EventObject
>;

export type log_array_address_array_EventFilter =
  TypedEventFilter<log_array_address_array_Event>;

export interface log_bytesEventObject {
  arg0: string;
}
export type log_bytesEvent = TypedEvent<[string], log_bytesEventObject>;

export type log_bytesEventFilter = TypedEventFilter<log_bytesEvent>;

export interface log_bytes32EventObject {
  arg0: string;
}
export type log_bytes32Event = TypedEvent<[string], log_bytes32EventObject>;

export type log_bytes32EventFilter = TypedEventFilter<log_bytes32Event>;

export interface log_intEventObject {
  arg0: BigNumber;
}
export type log_intEvent = TypedEvent<[BigNumber], log_intEventObject>;

export type log_intEventFilter = TypedEventFilter<log_intEvent>;

export interface log_named_addressEventObject {
  key: string;
  val: string;
}
export type log_named_addressEvent = TypedEvent<
  [string, string],
  log_named_addressEventObject
>;

export type log_named_addressEventFilter =
  TypedEventFilter<log_named_addressEvent>;

export interface log_named_array_string_uint256_array_EventObject {
  key: string;
  val: BigNumber[];
}
export type log_named_array_string_uint256_array_Event = TypedEvent<
  [string, BigNumber[]],
  log_named_array_string_uint256_array_EventObject
>;

export type log_named_array_string_uint256_array_EventFilter =
  TypedEventFilter<log_named_array_string_uint256_array_Event>;

export interface log_named_array_string_int256_array_EventObject {
  key: string;
  val: BigNumber[];
}
export type log_named_array_string_int256_array_Event = TypedEvent<
  [string, BigNumber[]],
  log_named_array_string_int256_array_EventObject
>;

export type log_named_array_string_int256_array_EventFilter =
  TypedEventFilter<log_named_array_string_int256_array_Event>;

export interface log_named_array_string_address_array_EventObject {
  key: string;
  val: string[];
}
export type log_named_array_string_address_array_Event = TypedEvent<
  [string, string[]],
  log_named_array_string_address_array_EventObject
>;

export type log_named_array_string_address_array_EventFilter =
  TypedEventFilter<log_named_array_string_address_array_Event>;

export interface log_named_bytesEventObject {
  key: string;
  val: string;
}
export type log_named_bytesEvent = TypedEvent<
  [string, string],
  log_named_bytesEventObject
>;

export type log_named_bytesEventFilter = TypedEventFilter<log_named_bytesEvent>;

export interface log_named_bytes32EventObject {
  key: string;
  val: string;
}
export type log_named_bytes32Event = TypedEvent<
  [string, string],
  log_named_bytes32EventObject
>;

export type log_named_bytes32EventFilter =
  TypedEventFilter<log_named_bytes32Event>;

export interface log_named_decimal_intEventObject {
  key: string;
  val: BigNumber;
  decimals: BigNumber;
}
export type log_named_decimal_intEvent = TypedEvent<
  [string, BigNumber, BigNumber],
  log_named_decimal_intEventObject
>;

export type log_named_decimal_intEventFilter =
  TypedEventFilter<log_named_decimal_intEvent>;

export interface log_named_decimal_uintEventObject {
  key: string;
  val: BigNumber;
  decimals: BigNumber;
}
export type log_named_decimal_uintEvent = TypedEvent<
  [string, BigNumber, BigNumber],
  log_named_decimal_uintEventObject
>;

export type log_named_decimal_uintEventFilter =
  TypedEventFilter<log_named_decimal_uintEvent>;

export interface log_named_intEventObject {
  key: string;
  val: BigNumber;
}
export type log_named_intEvent = TypedEvent<
  [string, BigNumber],
  log_named_intEventObject
>;

export type log_named_intEventFilter = TypedEventFilter<log_named_intEvent>;

export interface log_named_stringEventObject {
  key: string;
  val: string;
}
export type log_named_stringEvent = TypedEvent<
  [string, string],
  log_named_stringEventObject
>;

export type log_named_stringEventFilter =
  TypedEventFilter<log_named_stringEvent>;

export interface log_named_uintEventObject {
  key: string;
  val: BigNumber;
}
export type log_named_uintEvent = TypedEvent<
  [string, BigNumber],
  log_named_uintEventObject
>;

export type log_named_uintEventFilter = TypedEventFilter<log_named_uintEvent>;

export interface log_stringEventObject {
  arg0: string;
}
export type log_stringEvent = TypedEvent<[string], log_stringEventObject>;

export type log_stringEventFilter = TypedEventFilter<log_stringEvent>;

export interface log_uintEventObject {
  arg0: BigNumber;
}
export type log_uintEvent = TypedEvent<[BigNumber], log_uintEventObject>;

export type log_uintEventFilter = TypedEventFilter<log_uintEvent>;

export interface logsEventObject {
  arg0: string;
}
export type logsEvent = TypedEvent<[string], logsEventObject>;

export type logsEventFilter = TypedEventFilter<logsEvent>;

export interface GatewayEVMZEVMTest extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: GatewayEVMZEVMTestInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    IS_TEST(overrides?: CallOverrides): Promise<[boolean]>;

    excludeArtifacts(
      overrides?: CallOverrides
    ): Promise<[string[]] & { excludedArtifacts_: string[] }>;

    excludeContracts(
      overrides?: CallOverrides
    ): Promise<[string[]] & { excludedContracts_: string[] }>;

    excludeSelectors(
      overrides?: CallOverrides
    ): Promise<
      [StdInvariant.FuzzSelectorStructOutput[]] & {
        excludedSelectors_: StdInvariant.FuzzSelectorStructOutput[];
      }
    >;

    excludeSenders(
      overrides?: CallOverrides
    ): Promise<[string[]] & { excludedSenders_: string[] }>;

    failed(overrides?: CallOverrides): Promise<[boolean]>;

    setUp(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    targetArtifactSelectors(
      overrides?: CallOverrides
    ): Promise<
      [StdInvariant.FuzzArtifactSelectorStructOutput[]] & {
        targetedArtifactSelectors_: StdInvariant.FuzzArtifactSelectorStructOutput[];
      }
    >;

    targetArtifacts(
      overrides?: CallOverrides
    ): Promise<[string[]] & { targetedArtifacts_: string[] }>;

    targetContracts(
      overrides?: CallOverrides
    ): Promise<[string[]] & { targetedContracts_: string[] }>;

    targetInterfaces(
      overrides?: CallOverrides
    ): Promise<
      [StdInvariant.FuzzInterfaceStructOutput[]] & {
        targetedInterfaces_: StdInvariant.FuzzInterfaceStructOutput[];
      }
    >;

    targetSelectors(
      overrides?: CallOverrides
    ): Promise<
      [StdInvariant.FuzzSelectorStructOutput[]] & {
        targetedSelectors_: StdInvariant.FuzzSelectorStructOutput[];
      }
    >;

    targetSenders(
      overrides?: CallOverrides
    ): Promise<[string[]] & { targetedSenders_: string[] }>;

    testCallReceiverEVMFromSenderZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    testCallReceiverEVMFromZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    testWithdrawAndCallReceiverEVMFromSenderZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    testWithdrawAndCallReceiverEVMFromZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  IS_TEST(overrides?: CallOverrides): Promise<boolean>;

  excludeArtifacts(overrides?: CallOverrides): Promise<string[]>;

  excludeContracts(overrides?: CallOverrides): Promise<string[]>;

  excludeSelectors(
    overrides?: CallOverrides
  ): Promise<StdInvariant.FuzzSelectorStructOutput[]>;

  excludeSenders(overrides?: CallOverrides): Promise<string[]>;

  failed(overrides?: CallOverrides): Promise<boolean>;

  setUp(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  targetArtifactSelectors(
    overrides?: CallOverrides
  ): Promise<StdInvariant.FuzzArtifactSelectorStructOutput[]>;

  targetArtifacts(overrides?: CallOverrides): Promise<string[]>;

  targetContracts(overrides?: CallOverrides): Promise<string[]>;

  targetInterfaces(
    overrides?: CallOverrides
  ): Promise<StdInvariant.FuzzInterfaceStructOutput[]>;

  targetSelectors(
    overrides?: CallOverrides
  ): Promise<StdInvariant.FuzzSelectorStructOutput[]>;

  targetSenders(overrides?: CallOverrides): Promise<string[]>;

  testCallReceiverEVMFromSenderZEVM(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  testCallReceiverEVMFromZEVM(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  testWithdrawAndCallReceiverEVMFromSenderZEVM(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  testWithdrawAndCallReceiverEVMFromZEVM(
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    IS_TEST(overrides?: CallOverrides): Promise<boolean>;

    excludeArtifacts(overrides?: CallOverrides): Promise<string[]>;

    excludeContracts(overrides?: CallOverrides): Promise<string[]>;

    excludeSelectors(
      overrides?: CallOverrides
    ): Promise<StdInvariant.FuzzSelectorStructOutput[]>;

    excludeSenders(overrides?: CallOverrides): Promise<string[]>;

    failed(overrides?: CallOverrides): Promise<boolean>;

    setUp(overrides?: CallOverrides): Promise<void>;

    targetArtifactSelectors(
      overrides?: CallOverrides
    ): Promise<StdInvariant.FuzzArtifactSelectorStructOutput[]>;

    targetArtifacts(overrides?: CallOverrides): Promise<string[]>;

    targetContracts(overrides?: CallOverrides): Promise<string[]>;

    targetInterfaces(
      overrides?: CallOverrides
    ): Promise<StdInvariant.FuzzInterfaceStructOutput[]>;

    targetSelectors(
      overrides?: CallOverrides
    ): Promise<StdInvariant.FuzzSelectorStructOutput[]>;

    targetSenders(overrides?: CallOverrides): Promise<string[]>;

    testCallReceiverEVMFromSenderZEVM(overrides?: CallOverrides): Promise<void>;

    testCallReceiverEVMFromZEVM(overrides?: CallOverrides): Promise<void>;

    testWithdrawAndCallReceiverEVMFromSenderZEVM(
      overrides?: CallOverrides
    ): Promise<void>;

    testWithdrawAndCallReceiverEVMFromZEVM(
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {
    "Call(address,bytes,bytes)"(
      sender?: PromiseOrValue<string> | null,
      receiver?: null,
      message?: null
    ): Call_address_bytes_bytes_EventFilter;
    "Call(address,address,bytes)"(
      sender?: PromiseOrValue<string> | null,
      receiver?: PromiseOrValue<string> | null,
      payload?: null
    ): Call_address_address_bytes_EventFilter;

    "Deposit(address,address,uint256,address,bytes)"(
      sender?: PromiseOrValue<string> | null,
      receiver?: PromiseOrValue<string> | null,
      amount?: null,
      asset?: null,
      payload?: null
    ): DepositEventFilter;
    Deposit(
      sender?: PromiseOrValue<string> | null,
      receiver?: PromiseOrValue<string> | null,
      amount?: null,
      asset?: null,
      payload?: null
    ): DepositEventFilter;

    "Executed(address,uint256,bytes)"(
      destination?: PromiseOrValue<string> | null,
      value?: null,
      data?: null
    ): ExecutedEventFilter;
    Executed(
      destination?: PromiseOrValue<string> | null,
      value?: null,
      data?: null
    ): ExecutedEventFilter;

    "ExecutedWithERC20(address,address,uint256,bytes)"(
      token?: PromiseOrValue<string> | null,
      to?: PromiseOrValue<string> | null,
      amount?: null,
      data?: null
    ): ExecutedWithERC20EventFilter;
    ExecutedWithERC20(
      token?: PromiseOrValue<string> | null,
      to?: PromiseOrValue<string> | null,
      amount?: null,
      data?: null
    ): ExecutedWithERC20EventFilter;

    "ReceivedERC20(address,uint256,address,address)"(
      sender?: null,
      amount?: null,
      token?: null,
      destination?: null
    ): ReceivedERC20EventFilter;
    ReceivedERC20(
      sender?: null,
      amount?: null,
      token?: null,
      destination?: null
    ): ReceivedERC20EventFilter;

    "ReceivedNoParams(address)"(sender?: null): ReceivedNoParamsEventFilter;
    ReceivedNoParams(sender?: null): ReceivedNoParamsEventFilter;

    "ReceivedNonPayable(address,string[],uint256[],bool)"(
      sender?: null,
      strs?: null,
      nums?: null,
      flag?: null
    ): ReceivedNonPayableEventFilter;
    ReceivedNonPayable(
      sender?: null,
      strs?: null,
      nums?: null,
      flag?: null
    ): ReceivedNonPayableEventFilter;

    "ReceivedPayable(address,uint256,string,uint256,bool)"(
      sender?: null,
      value?: null,
      str?: null,
      num?: null,
      flag?: null
    ): ReceivedPayableEventFilter;
    ReceivedPayable(
      sender?: null,
      value?: null,
      str?: null,
      num?: null,
      flag?: null
    ): ReceivedPayableEventFilter;

    "Withdrawal(address,address,bytes,uint256,uint256,uint256,bytes)"(
      from?: PromiseOrValue<string> | null,
      zrc20?: null,
      to?: null,
      value?: null,
      gasfee?: null,
      protocolFlatFee?: null,
      message?: null
    ): WithdrawalEventFilter;
    Withdrawal(
      from?: PromiseOrValue<string> | null,
      zrc20?: null,
      to?: null,
      value?: null,
      gasfee?: null,
      protocolFlatFee?: null,
      message?: null
    ): WithdrawalEventFilter;

    "log(string)"(arg0?: null): logEventFilter;
    log(arg0?: null): logEventFilter;

    "log_address(address)"(arg0?: null): log_addressEventFilter;
    log_address(arg0?: null): log_addressEventFilter;

    "log_array(uint256[])"(val?: null): log_array_uint256_array_EventFilter;
    "log_array(int256[])"(val?: null): log_array_int256_array_EventFilter;
    "log_array(address[])"(val?: null): log_array_address_array_EventFilter;

    "log_bytes(bytes)"(arg0?: null): log_bytesEventFilter;
    log_bytes(arg0?: null): log_bytesEventFilter;

    "log_bytes32(bytes32)"(arg0?: null): log_bytes32EventFilter;
    log_bytes32(arg0?: null): log_bytes32EventFilter;

    "log_int(int256)"(arg0?: null): log_intEventFilter;
    log_int(arg0?: null): log_intEventFilter;

    "log_named_address(string,address)"(
      key?: null,
      val?: null
    ): log_named_addressEventFilter;
    log_named_address(key?: null, val?: null): log_named_addressEventFilter;

    "log_named_array(string,uint256[])"(
      key?: null,
      val?: null
    ): log_named_array_string_uint256_array_EventFilter;
    "log_named_array(string,int256[])"(
      key?: null,
      val?: null
    ): log_named_array_string_int256_array_EventFilter;
    "log_named_array(string,address[])"(
      key?: null,
      val?: null
    ): log_named_array_string_address_array_EventFilter;

    "log_named_bytes(string,bytes)"(
      key?: null,
      val?: null
    ): log_named_bytesEventFilter;
    log_named_bytes(key?: null, val?: null): log_named_bytesEventFilter;

    "log_named_bytes32(string,bytes32)"(
      key?: null,
      val?: null
    ): log_named_bytes32EventFilter;
    log_named_bytes32(key?: null, val?: null): log_named_bytes32EventFilter;

    "log_named_decimal_int(string,int256,uint256)"(
      key?: null,
      val?: null,
      decimals?: null
    ): log_named_decimal_intEventFilter;
    log_named_decimal_int(
      key?: null,
      val?: null,
      decimals?: null
    ): log_named_decimal_intEventFilter;

    "log_named_decimal_uint(string,uint256,uint256)"(
      key?: null,
      val?: null,
      decimals?: null
    ): log_named_decimal_uintEventFilter;
    log_named_decimal_uint(
      key?: null,
      val?: null,
      decimals?: null
    ): log_named_decimal_uintEventFilter;

    "log_named_int(string,int256)"(
      key?: null,
      val?: null
    ): log_named_intEventFilter;
    log_named_int(key?: null, val?: null): log_named_intEventFilter;

    "log_named_string(string,string)"(
      key?: null,
      val?: null
    ): log_named_stringEventFilter;
    log_named_string(key?: null, val?: null): log_named_stringEventFilter;

    "log_named_uint(string,uint256)"(
      key?: null,
      val?: null
    ): log_named_uintEventFilter;
    log_named_uint(key?: null, val?: null): log_named_uintEventFilter;

    "log_string(string)"(arg0?: null): log_stringEventFilter;
    log_string(arg0?: null): log_stringEventFilter;

    "log_uint(uint256)"(arg0?: null): log_uintEventFilter;
    log_uint(arg0?: null): log_uintEventFilter;

    "logs(bytes)"(arg0?: null): logsEventFilter;
    logs(arg0?: null): logsEventFilter;
  };

  estimateGas: {
    IS_TEST(overrides?: CallOverrides): Promise<BigNumber>;

    excludeArtifacts(overrides?: CallOverrides): Promise<BigNumber>;

    excludeContracts(overrides?: CallOverrides): Promise<BigNumber>;

    excludeSelectors(overrides?: CallOverrides): Promise<BigNumber>;

    excludeSenders(overrides?: CallOverrides): Promise<BigNumber>;

    failed(overrides?: CallOverrides): Promise<BigNumber>;

    setUp(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    targetArtifactSelectors(overrides?: CallOverrides): Promise<BigNumber>;

    targetArtifacts(overrides?: CallOverrides): Promise<BigNumber>;

    targetContracts(overrides?: CallOverrides): Promise<BigNumber>;

    targetInterfaces(overrides?: CallOverrides): Promise<BigNumber>;

    targetSelectors(overrides?: CallOverrides): Promise<BigNumber>;

    targetSenders(overrides?: CallOverrides): Promise<BigNumber>;

    testCallReceiverEVMFromSenderZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    testCallReceiverEVMFromZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    testWithdrawAndCallReceiverEVMFromSenderZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    testWithdrawAndCallReceiverEVMFromZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    IS_TEST(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    excludeArtifacts(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    excludeContracts(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    excludeSelectors(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    excludeSenders(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    failed(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    setUp(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    targetArtifactSelectors(
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    targetArtifacts(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    targetContracts(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    targetInterfaces(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    targetSelectors(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    targetSenders(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    testCallReceiverEVMFromSenderZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    testCallReceiverEVMFromZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    testWithdrawAndCallReceiverEVMFromSenderZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    testWithdrawAndCallReceiverEVMFromZEVM(
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}

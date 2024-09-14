import subprocess
import re
import os
import sys

from config import Config, MAINNET, TESTNET


class Deploy:

    def __init__(self, rpc, foundry_script_path):
        self.rpc = rpc
        self.foundry_script_path = foundry_script_path
        self.verifier_address = None
        self.plugman_address = None
        self.shuttle_address = None

    def deploy(self):
        self.deploy_verifier()
        self.deploy_plugman()
        self.deploy_shuttle_ethereum()
        self.set_up_shuttle_address()
        self.deploy_shuttle_polygon()
        self.set_up_shuttle_address()

    def deploy_verifier(self):
        try:
            script_path = f'{self.foundry_script_path}/DeployMintVerifier.s.sol'
            result = subprocess.run(
                ['forge', 'script', script_path, '--fork-url', self.rpc, '--broadcast'],
                capture_output=True,
                text=True)

            print(f'[Deploy - deploy_verifier] result is {result.stdout}')
            print(f'error output is {result.stderr}')

            match = re.search(r"Verifier contract is deployed at (\w+)", result.stdout)
            if match:
                self.verifier_address = match.group(1)
                os.environ['VERIFIER_ADDRESS'] = self.verifier_address
                print(f'[Deploy - deploy_verifier] verifier address is {self.verifier_address}')
            else:
                print(f'[Deploy - deploy_verifier] error: address is not found')
                exit(1)
        except subprocess.CalledProcessError as e:
            print(f'[Deploy - deploy_verifier] subprocess call error: {e}')
            exit(1)

    def deploy_plugman(self):
        try:
            script_path = f'{self.foundry_script_path}/DeployPlugman.s.sol'
            result = subprocess.run(
                ['forge', 'script', script_path, '--fork-url', self.rpc, '--broadcast'],
                capture_output=True,
                text=True)

            print(f'[Deploy - deploy_plugman] result is {result.stdout}')
            print(f'error output is {result.stderr}')

            match = re.search(r"Plugman contract is deployed at (\w+)", result.stdout)
            if match:
                self.plugman_address = match.group(1)
                os.environ['PLUGMAN_ADDRESS'] = self.plugman_address
                print(f'[Deploy - deploy_plugman] plugman address is {self.plugman_address}')
            else:
                print(f'[Deploy - deploy_plugman] error: address is not found')
                exit(1)
        except subprocess.CalledProcessError as e:
            print(f'[Deploy - deploy_plugman] subprocess call error: {e}')
            exit(1)

    def deploy_shuttle_ethereum(self):
        try:
            permit_chain_id = os.getenv("PERMIT_ETHEREUM_CHAIN_ID")
            crossing_wl_price = os.getenv("ETHEREUM_WL_PRICE")
            crossing_public_sale_price = os.getenv("ETHEREUM_PUBLIC_PRICE")
            os.environ["PERMIT_CHAIN_ID"] = permit_chain_id
            os.environ["CROSSING_WL_PRICE"] = crossing_wl_price
            os.environ["CROSSING_PUBLIC_SALE_PRICE"] = crossing_public_sale_price
            script_path = f'{self.foundry_script_path}/DeployShuttleMachine.s.sol'
            result = subprocess.run(
                ['forge', 'script', script_path, '--fork-url', self.rpc, '--broadcast'],
                capture_output=True,
                text=True)
            print(f'[Deploy - deploy_shuttle] result is {result.stdout}')

            match = re.search(r"Shuttle machine contract is deployed at (\w+)", result.stdout)
            if match:
                self.shuttle_address = match.group(1)
                os.environ['SHUTTLE_MACHINE_ADDRESS'] = self.shuttle_address
                print(f'[Deploy - deploy_shuttle] shuttle address is {self.shuttle_address}')
            else:
                print(f'[Deploy - deploy_shuttle] error: address is not found')
                exit(1)
        except subprocess.CalledProcessError as e:
            print(f'[Deploy - deploy_shuttle] subprocess call error: {e}')
            exit(1)

    def deploy_shuttle_polygon(self):
        try:
            permit_chain_id = os.getenv("PERMIT_POLYGON_CHAIN_ID")
            crossing_wl_price = os.getenv("POLYGON_WL_PRICE")
            crossing_public_sale_price = os.getenv("POLYGON_PUBLIC_PRICE")
            os.environ["PERMIT_CHAIN_ID"] = permit_chain_id
            os.environ["CROSSING_WL_PRICE"] = crossing_wl_price
            os.environ["CROSSING_PUBLIC_SALE_PRICE"] = crossing_public_sale_price
            script_path = f'{self.foundry_script_path}/DeployShuttleMachine.s.sol'
            result = subprocess.run(
                ['forge', 'script', script_path, '--fork-url', self.rpc, '--broadcast'],
                capture_output=True,
                text=True)
            print(f'[Deploy - deploy_shuttle] result is {result.stdout}')

            match = re.search(r"Shuttle machine contract is deployed at (\w+)", result.stdout)
            if match:
                self.shuttle_address = match.group(1)
                os.environ['SHUTTLE_MACHINE_ADDRESS'] = self.shuttle_address
                print(f'[Deploy - deploy_shuttle] shuttle address is {self.shuttle_address}')
            else:
                print(f'[Deploy - deploy_shuttle] error: address is not found')
                exit(1)
        except subprocess.CalledProcessError as e:
            print(f'[Deploy - deploy_shuttle] subprocess call error: {e}')
            exit(1)

    def set_up_shuttle_address(self):
        try:
            script_path = f'{self.foundry_script_path}/SetUpShuttleAddress.s.sol'
            result = subprocess.run(
                ['forge', 'script', script_path, '--fork-url', self.rpc, '--broadcast'],
                capture_output=True,
                text=True)
            print(f'[Deploy - set_up_shuttle_address] result is {result.stdout}')

            match = re.search("Set shuttle machine contract address on plugman", result.stdout)
            if match:
                print(f'[Deploy - set_up_shuttle_address] set up shuttle address successfully')
            else:
                print(f'[Deploy - set_up_shuttle_address] error: set up shuttle address failed')
                exit(1)
        except subprocess.CalledProcessError as e:
            print(f'[Deploy - set_up_shuttle_address] subprocess call error: {e}')
            exit(1)


def set_mainnet_env():
    os.environ["ZETA_SYSTEM_CONTRACT"] = "0x91d18e54DAf4F677cB28167158d6dd21F6aB3921"
    os.environ["TREASURY_ADDRESS"] = "need_to_be_filled"
    os.environ["PERMIT_ETHEREUM_CHAIN_ID"] = "1"
    os.environ["PERMIT_POLYGON_CHAIN_ID"] = "137"
    os.environ["ETHEREUM_WL_PRICE"] = "need_to_be_filled"
    os.environ["ETHEREUM_PUBLIC_PRICE"] = "need_to_be_filled"
    os.environ["POLYGON_WL_PRICE"] = "need_to_be_filled"
    os.environ["POLYGON_PUBLIC_PRICE"] = "need_to_be_filled"
    os.environ["METADATA_CREATOR"] = "need_to_be_filled"
    os.environ["BATCH_MAX"] = "999"


def set_testnet_env():
    os.environ["ZETA_SYSTEM_CONTRACT"] = "0xEdf1c3275d13489aCdC6cD6eD246E72458B8795B"
    os.environ["TREASURY_ADDRESS"] = "0x6F881627057b37B12163118e09F7e7901096Ccea"
    os.environ["PERMIT_ETHEREUM_CHAIN_ID"] = "11155111"
    os.environ["PERMIT_POLYGON_CHAIN_ID"] = "80002"
    os.environ["ETHEREUM_WL_PRICE"] = "2900000000000000"
    os.environ["ETHEREUM_PUBLIC_PRICE"] = "4900000000000000"
    os.environ["POLYGON_WL_PRICE"] = "6900000000000000"
    os.environ["POLYGON_PUBLIC_PRICE"] = "8900000000000000"
    os.environ["METADATA_CREATOR"] = "0x0d228F021B1F07e02642737Fbd4d46475A603eab"
    os.environ["BATCH_MAX"] = "999"


def set_localnet_env():
    os.environ["ZETA_SYSTEM_CONTRACT"] = "0x610178dA211FEF7D417bC0e6FeD39F05609AD788"
    os.environ["TREASURY_ADDRESS"] = "0x6F881627057b37B12163118e09F7e7901096Ccea"
    os.environ["PERMIT_ETHEREUM_CHAIN_ID"] = "31337"
    os.environ["PERMIT_POLYGON_CHAIN_ID"] = "31337"
    os.environ["ETHEREUM_WL_PRICE"] = "2900000000000000"
    os.environ["ETHEREUM_PUBLIC_PRICE"] = "4900000000000000"
    os.environ["POLYGON_WL_PRICE"] = "6900000000000000"
    os.environ["POLYGON_PUBLIC_PRICE"] = "8900000000000000"
    os.environ["METADATA_CREATOR"] = "0x0d228F021B1F07e02642737Fbd4d46475A603eab"
    os.environ["BATCH_MAX"] = "999"


if __name__ == '__main__':
    if Config.network == MAINNET:
        set_mainnet_env()
    elif Config.network == TESTNET:
        set_testnet_env()
    else:
        set_localnet_env()
    if os.getenv("TREASURY_ADDRESS") == "need_to_be_filled":
        print("Failed to create env, some value need to be changed")
        sys.exit(1)
    deploy = Deploy(Config.rpc, Config.foundry_script_path)
    deploy.deploy()

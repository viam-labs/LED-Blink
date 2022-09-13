import asyncio

from viam.robot.client import RobotClient
from viam.rpc.dial import Credentials, DialOptions
from viam.components.board import Board


async def connect():
    ADDRESS_SECRET = "[ADD YOUR ROBOT ADDRESS HERE. YOU CAN FIND THIS ON THE CONNECT TAB]"
    PAYLOAD_SECRET = "[PLEASE ADD YOUR SECRET HERE. YOU CAN FIND THIS ON THE CONNECT TAB]"
    creds = Credentials(
        type='robot-location-secret',
        payload=PAYLOAD_SECRET)
    opts = RobotClient.Options(
        refresh_interval=0,
        dial_options=DialOptions(credentials=creds)
    )
    return await RobotClient.at_address(ADDRESS_SECRET, opts)


async def main():
    robot = await connect()

    print('Resources:')
    print(robot.resource_names)

    local = Board.from_robot(robot, 'local')
    led = await local.gpio_pin_by_name('8')

    while (True):
        # When True, sets the LED pin to high or on.
        await led.set(True)
        print('LED is on')
        await asyncio.sleep(1)

        # When False, sets the pin to low or off.
        await led.set(False)
        print('LED is off')
        await asyncio.sleep(1)

    await robot.close()

if __name__ == '__main__':
    asyncio.run(main())

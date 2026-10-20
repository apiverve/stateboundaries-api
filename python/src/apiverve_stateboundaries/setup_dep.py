from setuptools import setup, find_packages

setup(
    name='apiverve_stateboundaries',
    version='1.2.2',
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        'requests',
        'setuptools'
    ],
    description='State Boundaries provides GeoJSON data and Census identifiers for US states using a two-letter postal code. Responses include the state name, FIPS code, and GEOID; paid plans add border polygon coordinates and land area measurements.',
    author='APIVerve',
    author_email='hello@apiverve.com',
    url='https://apiverve.com/marketplace/stateboundaries?utm_source=pypi&utm_medium=referral&utm_content=homepage',
    classifiers=[
        'Programming Language :: Python :: 3',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
)

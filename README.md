# Dependency Injection in Go

This project shows how to solve a a dependency injection problem in Go.

Keywords: dependency injection, inject, injector, gin

## Introduction

Dependency injection seems to be a key factor to write a DRI code, especially when a developer does not want to care about passing a service through few layers, always accompanied by a long list of constructor arguments. This boilerplate usually makes the code unreadable and unnecessarily longer than the crucial business logic.

The way I figured this out, is presented in a "proof of concept" simple project. I tried to eliminate as much of not related to DI source code, thus sometimes you may wonder if this simplicity really needs a DI pattern.

## User stories

* There is a library in which a user can borrow a book.
* The user has name
* The book has a title
* The user can borrow a book.

## Solution Design

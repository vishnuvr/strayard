#
# This is the base image from which all StraYard images inherit.
#
# The standard name for this image is strayard/base
#
FROM centos:centos7

# components from EPEL must be installed in a separate yum install step
RUN yum install -y git tar wget socat hostname epel-release && yum clean all